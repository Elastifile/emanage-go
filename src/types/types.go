package types

import (
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	log15 "gopkg.in/inconshreveable/log15.v2"

	filestore_types "github.com/elastifile/emanage-go/src/filestore/types"
	logging_config "github.com/elastifile/emanage-go/src/logging/config"
	"github.com/elastifile/emanage-go/src/optional"
)

// A Host is used to represent a host by its hostname or IP address.
type Host string

func (h Host) String() string { return string(h) }

const (
	HostEmpty Host = ""
)

func Hosts(hosts []string) []Host {
	result := make([]Host, len(hosts))
	for i, h := range hosts {
		result[i] = Host(h)
	}
	return result
}

func HostStrings(hosts []Host) []string {
	result := make([]string, len(hosts))
	for i, h := range hosts {
		result[i] = string(h)
	}
	return result
}

// A Tesla JobID
type JobID string

var randID = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewID() JobID {
	// Should we support a longer ID? We use just 6 bytes to be nice and short.
	// Since we prevent collisions in AddNew, this should not be a problem in practice.
	data := make([]byte, 6)
	for i := 0; i < len(data); i++ {
		data[i] = byte(randID.Intn(256))
	}
	return JobID(hex.EncodeToString(data))
}

// A Tesla Job status
//go:generate stringer -type=JobStatus

type JobStatus int

const (
	JobStatusSuccess JobStatus = iota
	JobStatusStarted
	JobStatusRunning
	JobStatusFailed
	JobStatusAborted
)

type ContainerStatus struct {
	ID   string
	Name string

	// Properties from docker.State
	Running    bool
	Paused     bool
	Restarting bool
	OOMKilled  bool
	Pid        int
	ExitCode   int
	Error      string
	StartedAt  time.Time
	FinishedAt time.Time
}

type ToolStatus struct {
	Containers map[Host][]*ContainerStatus
}

func (ts *ToolStatus) IsRunning() bool {
	for _, cs := range ts.Containers {
		for _, c := range cs {
			if c.Running {
				return true
			}
		}
	}
	return false
}

func (ts *ToolStatus) HasFailures() bool {
	for _, cs := range ts.Containers {
		for _, c := range cs {
			if c.OOMKilled || c.ExitCode != 0 || c.Error != "" {
				return true
			}
		}
	}
	return false
}

type Image struct {
	registry   string
	port       string
	repository string
	tag        string
}

func (im Image) Registry() string { return im.registry }

func (im Image) Port() string { return im.port }

func (im Image) Tag() string { return im.tag }

func (im Image) Repository() string { return im.repository }

func (im Image) Name() string {
	var port string
	if im.port != "" {
		port = ":" + im.port
	}
	return im.registry + port + "/" + im.repository
}

func (im Image) String() string {
	return im.Name() + ":" + im.tag
}

func (im Image) ShortName() string {
	parts := strings.Split(im.repository, "/")
	return parts[len(parts)-1]
}

func (im *Image) SetTag(tag string) {
	im.tag = tag
}

func (im *Image) UnmarshalJSON(data []byte) error {
	*im = ParseImage(string(data[1: len(data)-1]))
	return nil
}

func (im *Image) MarshalJSON() ([]byte, error) {
	return []byte("\"" + im.String() + "\""), nil
}

func ParseImage(data string) Image {
	parts := strings.Split(string(data), "/")
	registryURL := append(strings.Split(parts[0], ":"), "")
	rest := strings.Join(parts[1:], "/")
	ending := append(strings.Split(rest, ":"), "")
	return Image{
		registry:   registryURL[0],
		port:       registryURL[1],
		repository: ending[0],
		tag:        ending[1],
	}
}

// A Context is passed between tesla components and preserves request properties along the path.
type Context struct {
	ID      uint32
	Timeout time.Duration
}

// A Container represents a Docker container running on a specific host.
type Container struct {
	Name      string
	ID        string
	Host      Host
	StartTime time.Time
}

type LogAggregator interface {
	fmt.Stringer
	ProcessLine(line string)
}

// A Job comprises one or more Containers running on the tesla cluster.
type Job struct {
	JobInfo
	Err  error
	Tool Tool
}

// JobInfo is the JSON-marshalable part of a Job.
type JobInfo struct {
	Name                     string // User-supplied name
	ID                       JobID  // Autogenerated unique ID
	Children                 []JobID
	StartTime                time.Time
	Config                   Config
	Status                   JobStatus
	Containers               []Container
	ToolResults              []ToolUnifiedResults
	ToolResultsFilesPatterns []string
}

type ToolUnifiedResults struct {
	Identifier     optional.String
	ToolName       optional.String
	Workload       optional.String
	RequestedIOps  optional.Int
	IOps           optional.Int
	AverageLatency optional.Float64 // in msec
	Bandwidth      optional.Float64 // in msec
	JobID          optional.String
}

// A File represents a file (possibly a template) that can be sent as part of a configuration.
type File struct {
	Name    string
	Content []byte // []byte and not string, since we want its content to be opaquely marshalable.
}

// ToolParams include all that is needed to run a tool.
type ToolParams struct {
	ToolName          ToolName
	ImageTag          string
	Args              []string
	Config            Config
	System            *System
	ConfigFilesBucket filestore_types.Bucket
	TargetLoaders     []string

	// Some jobs are started before the tool is created
	JobID      JobID
	Identifier string
	ParentID   JobID

	// This is set by `tesla run` when run from Jenkins
	BuildNumber string
	Uploads     []string
}

func (tp *ToolParams) FilerSessionPath() string {
	if !tp.Config.Logging.Filer.Enabled {
		return ""
	}
	return filepath.Join(tp.Config.Logging.SessionPath(), "jobs", string(tp.JobID), string(tp.ToolName))
}

type FilesByHost map[Host][]*NamedReader

type NamedReader struct {
	Name   string
	reader io.Reader
}

func NewNamedReader(name string, r io.Reader) *NamedReader {
	return &NamedReader{
		Name:   name,
		reader: r,
	}
}

type ByNamedReaders []*NamedReader

func (nrs ByNamedReaders) Len() int {
	return len(nrs)
}

func (nrs ByNamedReaders) Swap(i, j int) {
	nrs[i], nrs[j] = nrs[j], nrs[i]
}

func (nrs ByNamedReaders) Less(i, j int) bool {
	return nrs[i].Name < nrs[j].Name
}

func (f *NamedReader) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *NamedReader) String() string {
	return f.Name
}

func (f *NamedReader) GoString() string {
	return fmt.Sprintf("%T{%s: %+v}", f, f.Name, f.reader)
}

type ResultOpts struct {
	FilesByHost FilesByHost
	Conf        Config
	JobID       JobID
}

type ToolProtocol interface {
	Start(*Context) error
	Wait(*Context) error
	Stop(*Context) error
	Cleanup(*Context) error
	Abort(*Context) error
}

type ToolProperties interface {
	Name() ToolName
	WaitFor() []Host
	Params() *ToolParams
	Logger() log15.Logger
	ErrorWatcher() chan error
	NewLogAggregator(*Container) LogAggregator
	Logs() map[Host]*NamedReader
}

// The general Tool interface: Any tool should implement this.
type Tool interface {
	ToolProtocol
	ToolProperties
}

// Functions of this kind are used in conjunction with
// tool_errors.FoldErrors.  They define a strategy for handling errors,
// typically received asynchronously.  For example, `BarfStrategy'
// will cause the calling site to respond at the first error received,
// while `AccumulateStrategy' will collect all errors.  Some packages
// implement more complex strategies, eg. erun.(*tool).WaitStrategy()
// will create a strategy for filtering some errors, while stopping on
// other errors.
type ErrorHandlingStrategy func([]error, error) ([]error, bool)

// External tools create and run tools in containers.
type ExternalProperties interface {
	ToolProperties
	GetImage() *Image
	GetMasterCommand() ([]string, error)
	GetSlaveCommand(Host) ([]string, error)
	GetDockerHostConfig() *docker.HostConfig
	Master() Host
	WaitStrategy(*Context) ErrorHandlingStrategy
}

// Logging tools include additional files in their results.
type LoggingProperties interface {
	ExternalProperties
	GetResults(*ResultOpts) error
	GetResultFilesPatterns() []string
}

type InternalClient interface {
	RequestStart(Host, *Context, JobID, interface{} /* config */) error
	RequestStop(Host, *Context, JobID) error
	RequestWait(Host, *Context, JobID) error
	RequestCleanup(*Context, JobID) error
}

// Internal tools run wholly inside tesla agents.
type InternalProperties interface {
	ToolProperties
	Config() interface{}
	Client() InternalClient
	Master() Host
}

type NewToolFunc func(params *ToolParams) (Tool, error)

type ContainerStartOpts struct {
	Host             Host
	Image            *Image
	Cmd              []string
	DataStorePath    string
	JobID            JobID
	DockerHostConfig *docker.HostConfig
}

type Deployment struct {
	Elab           Elab
	System         System
	Elfs           Elfs
	Emanage        Emanage
	Vcenter        VCenter
	Jenkins        Jenkins
	OVAinfo        OvaDeployment
	DeploymentData ElabDeployData
	Logging        logging_config.Config
	Cloud          CloudConnect
	ReportOutputs  Outputs
	Telegraf       Telegraf
}

func (d Deployment) Loaders() (loaders []Host) {
	for _, ip := range d.System.Elab.LoaderIps() {
		loaders = append(loaders, Host(ip))
	}
	return loaders
}

/*
   {
       "host": "h420.lab.il.elastifile.com",
       "hostname": null,
       "ip_address": null,
       "networks": [],
       "role": null,
       "state": "off",
       "type": "virtual",
       "vm_name": "ReplicationService02"
     },

*/

type PushPullOptions struct {
	Registry    string `default:"registry.il.elastifile.com" doc:"Docker registry to use"`
	Username    string `doc:"Docker registry username for pushing/pulling images"`
	Branch      string `doc:"Git branch from which the images were built, used to determine image tags"`
	NoPull      bool   `doc:"Don't pull (or push) from the registry; user local images only"`
	UsePersonal bool   `doc:"Use the user's personal tesla images instead of the official ones"`
	ImageTag    string `doc:"pull using this tag instead of digest"`
}

type CloudStorageType string

const (
	LocalStorageType      CloudStorageType = "local"
	PersistentStorageType CloudStorageType = "persistent"
	DisklessStorageType   CloudStorageType = "diskless"
)

type ElfsVersion string

const (
	ElfsVersion2_7   ElfsVersion = "2.7.0"
	ElfsVersion2_5_1 ElfsVersion = "2.5.1"
	ElfsVersion2_5_0 ElfsVersion = "2.5.0"
	ElfsVersion2_0   ElfsVersion = "2.0.0"
)