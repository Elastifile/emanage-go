package fstool_config

import (
	"fmt"
	"strings"

	"github.com/go-errors/errors"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/koding/multiconfig"
	yaml "gopkg.in/yaml.v2"

	"github.com/elastifile/emanage-go/src/helputils"
	"github.com/elastifile/emanage-go/src/nfs/sunrpc/nfsx"
	"github.com/elastifile/emanage-go/src/size"
)

type Config struct {
	Method       string           `yaml:"Method"`
	TreeCreate   TreeCreateOpts   `yaml:"TreeCreate"`
	TreeUpdate   TreeUpdateOpts   `yaml:"TreeUpdate"`
	TreeDelete   TreeDeleteOpts   `yaml:"TreeDelete"`
	TreeDiff     TreeDiffOpts     `yaml:"TreeDiff"`
	FillCapacity FillCapacityOpts `yaml:"FillCapacity"`
	NfsClient    NfsClientOpts    `yaml:"NfsClient"`
}

func NewConfig() (conf Config) {
	tagLoader := &multiconfig.TagLoader{}
	err := tagLoader.Load(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func (conf *Config) String() string {
	body, _ := yaml.Marshal(*conf)
	return helputils.YamlToOneLine(string(body), 2)
}

func (conf *Config) ToCmd() []string {
	return strings.Split(conf.String(), " ")
}

type NfsClientOpts struct {
	Frontend       string `yaml:"Frontend"`
	Port           int    `yaml:"Port"`
	Mountport      int    `yaml:"Mountort"`
	Export         string `yaml:"Export"`
	Uid            int    `yaml:"Uid"            default:"17000"`
	Gid            int    `yaml:"Gid"            default:"18000"`
	NumConnections int    `yaml:"NumConnections" default:"0"      doc:"'0' means CPU count"`
	Timeout        int    `yaml:"Timeout"        default:"30"`
	Directory      string `yaml:"Directory"      default:"topdir"`
	Loader         string `yaml:"Loader"`
}

func NewNfsClient() (conf NfsClientOpts) {
	tagLoader := &multiconfig.TagLoader{}
	err := tagLoader.Load(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}

func (opts *NfsClientOpts) Update(upd NfsClientOpts) {
	if upd.Frontend != "" {
		opts.Frontend = upd.Frontend
	}
	if upd.Port != 0 {
		opts.Port = upd.Port
	}
	if upd.Mountport != 0 {
		opts.Mountport = upd.Mountport
	}
	if upd.Export != "" {
		opts.Export = upd.Export
	}
	if upd.Uid != 0 {
		opts.Uid = upd.Uid
	}
	if upd.Gid != 0 {
		opts.Gid = upd.Gid
	}
	if upd.NumConnections != 0 {
		opts.NumConnections = upd.NumConnections
	}
	if upd.Timeout != 0 {
		opts.Timeout = upd.Timeout
	}
	if upd.Directory != "" {
		opts.Directory = upd.Directory
	}
	if upd.Loader != "" {
		opts.Loader = upd.Loader
	}
}

type FillCapacityOpts struct {
	Goal       float64 `yaml:"Goal"       default:"1.0"`
	NumThreads int     `yaml:"NumThreads" default:"20"`
	EmanageIp  string  `yaml:"EmanageIp"`
	Verify     bool    `yaml:"Verify"`
}

func NewFillCapacity(opts *FillCapacityOpts) Config {
	conf := NewConfig()
	conf.Method = "FillCapacity"

	if opts != nil {
		conf.FillCapacity.Update(*opts)
	}

	return conf
}

func (opts *FillCapacityOpts) Update(upd FillCapacityOpts) {
	if upd.Goal != 0 {
		opts.Goal = upd.Goal
	}
	if upd.NumThreads != 0 {
		opts.NumThreads = upd.NumThreads
	}
	if upd.EmanageIp != "" {
		opts.EmanageIp = upd.EmanageIp
	}
	if upd.Verify {
		opts.Verify = upd.Verify
	}
}

type TreeCreateOpts struct {
	FileCount int           `yaml:"FileCount"`
	FileSize  int64         `yaml:"FileSize"`
	LinkCount int           `yaml:"LinkCount"             doc:"number of links upon created files"`
	NodeTypes NodeTypeNames `yaml:"NodeTypes"             doc:"CSV of type names"`
	NodeCount int           `yaml:"NodeCount"             doc:"number of nodes per types = names * count"`
	WriteData []byte        `yaml:"WriteData"             doc:"data to write, repeatedly - empty means random"`
	DirDepth  int           `yaml:"DirDepth"  default:"4" doc:"depth of tree"`
	ChildDirs int           `yaml:"ChildDirs" default:"4" doc:"number of child-dirs per dir"`
}

func NewTreeCreate(opts *TreeCreateOpts) Config {
	conf := NewConfig()
	conf.Method = "TreeCreate"

	if opts != nil {
		conf.TreeCreate.Update(*opts)
	}

	return conf
}

func (opts *TreeCreateOpts) Update(upd TreeCreateOpts) {
	if upd.FileCount != 0 {
		opts.FileCount = upd.FileCount
	}
	if upd.FileSize != 0 {
		opts.FileSize = upd.FileSize
	}
	if upd.LinkCount != 0 {
		opts.LinkCount = upd.LinkCount
	}
	if upd.NodeTypes != "" {
		opts.NodeTypes = upd.NodeTypes
	}
	if upd.NodeCount != 0 {
		opts.NodeCount = upd.NodeCount
	}
	if len(upd.WriteData) > 0 {
		opts.WriteData = upd.WriteData
	}
	if upd.ChildDirs != 0 {
		opts.ChildDirs = upd.ChildDirs
	}
	if upd.DirDepth != 0 {
		opts.DirDepth = upd.DirDepth
	}
}

type TreeUpdateOpts struct {
	Select       TreeSelectOpts `yaml:"Select"`
	WriteSize    int64          `yaml:"WriteSize"    default:"0"`
	WriteOffset  int64          `yaml:"WriteOffset"  default:"0"`
	WriteData    []byte         `yaml:"WriteData"                  doc:"data to write, repeatedly - empty means random"`
	WriteRepeat  int            `yaml:"WriteRepeat"`
	TruncateSize int64          `yaml:"TruncateSize" default:"-1"`
	FileAttrs    FileAttrOpts   `yaml:"FileAttrs"`
	Move         string         `yaml:"Move"`
}

func NewTreeUpdate(opts *TreeUpdateOpts) Config {
	conf := NewConfig()
	conf.Method = "TreeUpdate"

	if opts != nil {
		conf.TreeUpdate.Update(*opts)
	}

	return conf
}

func (opts *TreeUpdateOpts) Update(upd TreeUpdateOpts) {
	if upd.WriteSize != 0 {
		opts.WriteSize = upd.WriteSize
	}
	if upd.WriteOffset != 0 {
		opts.WriteOffset = upd.WriteOffset
	}
	if len(upd.WriteData) > 0 {
		opts.WriteData = upd.WriteData
	}
	if upd.WriteRepeat != 0 {
		opts.WriteRepeat = upd.WriteRepeat
	}
	if upd.TruncateSize != 0 {
		opts.TruncateSize = upd.TruncateSize
	}
	if upd.Move != "" {
		opts.Move = upd.Move
	}

	opts.Select.Update(upd.Select)
	opts.FileAttrs.Update(upd.FileAttrs)
}

type TreeDeleteOpts struct {
	Select TreeSelectOpts `yaml:"Select"`
}

func NewTreeDelete(opts *TreeDeleteOpts) Config {
	conf := NewConfig()
	conf.Method = "TreeDelete"

	if opts != nil {
		conf.TreeDelete.Update(*opts)
	}

	return conf
}

func (opts *TreeDeleteOpts) Update(upd TreeDeleteOpts) {
	opts.Select.Update(upd.Select)
}

type TreeDiffOpts struct {
	OtherNfsClient NfsClientOpts `yaml:"OtherNfsClient"`
	Ctime          bool          `yaml:"Ctime"`
}

func NewTreeDiff(opts *TreeDiffOpts) Config {
	conf := NewConfig()
	conf.Method = "TreeDiff"
	conf.TreeDiff.OtherNfsClient = conf.NfsClient
	conf.TreeDiff.OtherNfsClient.Frontend = "127.0.0.1"

	if opts != nil {
		conf.TreeDiff.Update(*opts)
	}

	return conf
}

func (opts *TreeDiffOpts) Update(upd TreeDiffOpts) {
	opts.OtherNfsClient.Update(upd.OtherNfsClient)
}

var fileTypeNames = strings.Split(fmt.Sprintf(",%s,%s,%s,%s,%s,%s,%s",
	nfsx.NF3REG,
	nfsx.NF3DIR,
	nfsx.NF3BLK,
	nfsx.NF3CHR,
	nfsx.NF3LNK,
	nfsx.NF3SOCK,
	nfsx.NF3FIFO,
), ",")

type NodeTypeNames string

func (ftyns NodeTypeNames) Verify() (err error) {
	for _, tyn := range strings.Split(string(ftyns), ",") {
		if !helputils.ContainsStr(fileTypeNames, tyn) {
			err = multierror.Append(err, errors.Errorf("invalid file type name: %s", tyn))
		}
	}
	return err
}

func (ftyns NodeTypeNames) Selected(typ nfsx.Ftype3) bool {
	return ftyns == "" ||
		helputils.ContainsStr(strings.Split(string(ftyns), ","), string(typ))
}

func (ftyns NodeTypeNames) Enum() (result []nfsx.Ftype3) {
	for _, tyn := range strings.Split(string(ftyns), ",") {
		result = append(result, nfsx.Ftype3Enum(tyn))
	}
	return result
}

type TreeSelectOpts struct {
	Percent   float64       `yaml:"Percent"   doc:"percent of files to update" default:"100.0"`
	Bigger    int64         `yaml:"Bigger"`
	Smaller   int64         `yaml:"Smaller"`
	NodeTypes NodeTypeNames `yaml:"NodeTypes" doc:"CSV of type names"`
	DirDepth  int           `yaml:"DirDepth"`
	Name      string        `yaml:"Name"      doc:"regular expression"`
}

func (opts *TreeSelectOpts) Update(upd TreeSelectOpts) {
	if upd.Percent != 0 {
		opts.Percent = upd.Percent
	}
	if upd.Bigger != 0 {
		opts.Bigger = upd.Bigger
	}
	if upd.Smaller != 0 {
		opts.Smaller = upd.Smaller
	}
	if len(upd.NodeTypes) > 0 {
		opts.NodeTypes = upd.NodeTypes
	}
	if upd.DirDepth != 0 {
		opts.DirDepth = upd.DirDepth
	}
	if upd.Name != "" {
		opts.Name = upd.Name
	}
}

type FileAttrOpts struct {
	Mode  int64 `yaml:"Mode"  default:"-1"`
	Uid   int64 `yaml:"Uid"   default:"-1"`
	Gid   int64 `yaml:"Gid"   default:"-1"`
	Size  int64 `yaml:"Size"  default:"-1"`
	Atime int64 `yaml:"Atime" default:"-1"`
	Mtime int64 `yaml:"Mtime" default:"-1"`
}

func (opts *FileAttrOpts) Update(upd FileAttrOpts) {
	if upd.Mode != 0 {
		opts.Mode = upd.Mode
	}
	if upd.Uid != 0 {
		opts.Uid = upd.Uid
	}
	if upd.Gid != 0 {
		opts.Gid = upd.Gid
	}
	if upd.Size != 0 {
		opts.Size = upd.Size
	}
	if upd.Atime != 0 {
		opts.Atime = upd.Atime
	}
	if upd.Mtime != 0 {
		opts.Mtime = upd.Mtime
	}
}

const maxFileSize = size.TiB

func (opts *NfsClientOpts) WorkDir() string {
	return opts.Directory + opts.Loader
}

func (opts *TreeCreateOpts) Verify() (err error) {
	if opts.FileCount < 0 {
		err = multierror.Append(err, errors.New("config invalid number of files"))
	}

	if opts.NodeCount < 0 {
		err = multierror.Append(err, errors.New("config invalid number of nodes"))
	}

	if opts.LinkCount < 0 {
		err = multierror.Append(err, errors.New("config invalid negaitive number of links"))
	}

	if opts.FileSize < 0 {
		err = multierror.Append(err, errors.New("config invalid negaitive file size"))
	} else if opts.FileSize > int64(maxFileSize) {
		err = multierror.Append(err, errors.New("config file size exceeds max of "+string(maxFileSize)))
	}

	if opts.ChildDirs < 0 {
		err = multierror.Append(err, errors.New("config invalid number of subdirs"))
	}

	if opts.DirDepth < 0 {
		err = multierror.Append(err, errors.New("config invalid depth of dirs"))
	}

	return err
}

func (opts *TreeUpdateOpts) Verify() (err error) {
	if e := opts.Select.Verify(); e != nil {
		err = multierror.Append(err, e)
	}

	if opts.WriteSize < 0 {
		err = multierror.Append(err, errors.New("invalid negaitive write size"))
	} else if opts.WriteSize > int64(maxFileSize) {
		err = multierror.Append(err, errors.New("write size exceeds max of "+string(maxFileSize)))
	}

	if opts.WriteOffset > int64(maxFileSize) {
		err = multierror.Append(err, errors.New("write offset exceeds max of "+string(maxFileSize)))
	}

	if opts.TruncateSize < 0 && opts.TruncateSize != -1 {
		err = multierror.Append(err, errors.New("invalid file negative selection smaller"))
	}

	if e := opts.FileAttrs.Verify(); e != nil {
		err = multierror.Append(err, e)
	}

	return err
}

func (opts *TreeDeleteOpts) Verify() (err error) {
	if e := opts.Select.Verify(); e != nil {
		err = multierror.Append(err, e)
	}

	return err
}

func (opts *TreeSelectOpts) Verify() (err error) {
	if opts.Percent < 0.0 || opts.Percent > 100.0 {
		err = multierror.Append(err, errors.New("invalid selection percent"))
	}

	if opts.Bigger < 0 {
		err = multierror.Append(err, errors.New("invalid negative selection bigger"))
	}

	if opts.Smaller < 0 {
		err = multierror.Append(err, errors.New("invalid negative selection smaller"))
	}

	if e := opts.NodeTypes.Verify(); e != nil {
		err = multierror.Append(err, e)
	}

	return err
}

func (opts *FileAttrOpts) Verify() (err error) {
	if opts.Mode < -1 {
		err = multierror.Append(err, errors.New("invalid negative file mode"))
	}

	if opts.Gid < -1 {
		err = multierror.Append(err, errors.New("invalid negative file gid"))
	}

	if opts.Uid < -1 {
		err = multierror.Append(err, errors.New("invalid negative file uid"))
	}

	if opts.Size < -1 {
		err = multierror.Append(err, errors.New("invalid negative file size"))
	} else if opts.Size > int64(maxFileSize) {
		err = multierror.Append(err, errors.New("file size exceeds max of "+string(maxFileSize)))
	}

	if opts.Mtime < -1 {
		err = multierror.Append(err, errors.New("invalid negative file mtime"))
	}

	if opts.Atime < -1 {
		err = multierror.Append(err, errors.New("invalid negative file atime"))
	}

	return err
}

func (opts *TreeCreateOpts) String() string {
	body, _ := yaml.Marshal(*opts)
	return helputils.YamlToOneLine(string(body), 2)
}

func (opts *TreeUpdateOpts) String() string {
	body, _ := yaml.Marshal(*opts)
	return helputils.YamlToOneLine(string(body), 2)
}

func (opts *TreeDeleteOpts) String() string {
	body, _ := yaml.Marshal(*opts)
	return helputils.YamlToOneLine(string(body), 2)
}

func (opts *TreeDiffOpts) String() string {
	body, _ := yaml.Marshal(*opts)
	return helputils.YamlToOneLine(string(body), 2)
}

func (opts *FillCapacityOpts) String() string {
	body, _ := yaml.Marshal(*opts)
	return helputils.YamlToOneLine(string(body), 2)
}

func (opts *NfsClientOpts) String() string {
	body, _ := yaml.Marshal(*opts)
	return helputils.YamlToOneLine(string(body), 2)
}

//go:generate stringer -type=FCRet

type FCRet int

const (
	FCRetSuccess        FCRet = 0
	FCRetErrorBadConfig FCRet = 1
	FCRetErrorNFSClient FCRet = 2
	FCRetErrorRootDir   FCRet = 3
	FCRetErrorFSInfo    FCRet = 4
	FCRetErrorWorkDir   FCRet = 5
	FCRetErrorTool      FCRet = 6
	FCRetErrorPanic     FCRet = 7
)
