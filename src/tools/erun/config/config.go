package erun_config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/go-errors/errors"

	"size"
)

const (
	erunLogsOnHost = "/var/log/erun.log"
)

type ErunFormatCmdOpts struct {
	Host             string
	HostIndex        int
	DataStorePath    string
	JobID            string
	Frontend         string
	Export           string
	ResultFile       string
	ResultsFullLog   string
	ExecPath         string
	FilerSessionPath string
	Conf             *Config
}

func (opts *ErunFormatCmdOpts) ErunDir() string {
	erunDir := opts.Conf.ErunDir
	if erunDir == "" {
		dirSuffix := opts.Conf.DirSuffix
		if dirSuffix == "" {
			dirSuffix = opts.JobID
		}
		return fmt.Sprintf("%s-%s", opts.Host, dirSuffix)
	}
	return erunDir
}

func (opts *ErunFormatCmdOpts) InBucket(file string) string {
	return filepath.Join(opts.DataStorePath, file)
}

func (opts *ErunFormatCmdOpts) LogSummary(tpl string, files ...string) string {
	var stored []interface{}
	for _, f := range files {
		stored = append(stored, opts.InBucket(f))
	}
	return fmt.Sprintf(tpl, stored...)
}

func (opts *ErunFormatCmdOpts) ParamFile() string {
	return filepath.Join(opts.DataStorePath, opts.Conf.ConfigFile)
}

func (opts *ErunFormatCmdOpts) ReportFile() string {
	return filepath.Join(opts.DataStorePath, opts.ResultFile)
}

func (opts *ErunFormatCmdOpts) Filesystem() string {
	return fmt.Sprintf("%s:%s", opts.Frontend, opts.Export)
}

const collectCoreDumpsScriptFmt = `
if [ -n "$(find /dev/shm -name erun*)" ]
then
	echo Copying erun core dumps to session dir...
	sudo mkdir -p %s
	sudo chmod -R g+w %s/../../..
	sudo cp /dev/shm/erun* %s/
	echo Deleting erun local core dumps...
	sudo rm -f /dev/shm/erun*
	echo Done.
fi`

func (opts *ErunFormatCmdOpts) CoreDumpsScript() string {
	result := "echo 'No job session path'"
	if opts.FilerSessionPath != "" {
		result = fmt.Sprintf(collectCoreDumpsScriptFmt,
			opts.FilerSessionPath,
			opts.FilerSessionPath,
			opts.FilerSessionPath,
		)
		result = " ; " + strings.Replace(result[1:], "\n", " ; ", -1)
		result = strings.Replace(result, "then ;", "then", -1)
	}
	return result
}

func (opts *ErunFormatCmdOpts) Combine(profile Profile, args ...string) string {
	result := strings.Join(
		append(
			[]string{opts.ExecPath, "--profile", profile.Name(), opts.Filesystem()},
			args...,
		),
		" ",
	)

	return result
}

type Profile interface {
	json.Unmarshaler
	Name() string
	ToFullCmd(*ErunFormatCmdOpts) string
	ToCmd() string
	HasResults() bool
}

func profiles() []Profile {
	return []Profile{
		&ProfileIO{},
		&ProfileMetabench{},
		&ProfileMetafuzzer{},
		&ProfileDevRegression{},
	}
}

// Examples:
//
// Profile: metafuzzer
// Cmd: --duration 30 --clients 3 --queue-size 5
// ErunDirSuffix: "metatest"
//
// Profile: io
// Cmd: --duration 30 --clients 1 --nr-files 1 --max-file-size 100M --queue-size 24 --readwrites 70 --data-payload --initial-write-phase
// ErunDirSuffix: ""
//
type Config struct {
	// when empty, generated randomly
	DirSuffix  string  `yaml:"DirSuffix"`
	ErunDir    string  `yaml:"ErunDir"`
	Profile    Profile `yaml:"Profile"`
	ConfigFile string  `yaml:"ConfigFile" tesla:"filename"`
}

func (c *Config) EnvMarshal(fieldName string) []string {
	return []string{
		fmt.Sprintf("%s_CMD=%s", fieldName, c.Profile.ToCmd()),
		fmt.Sprintf("%s_PROFILE=%s", fieldName, c.Profile.Name()),
		fmt.Sprintf("%s_CONFIG_FILE=%s", fieldName, c.ConfigFile),
		fmt.Sprintf("%s_DIR_SUFFIX=%s", fieldName, c.DirSuffix),
		fmt.Sprintf("%s_ERUN_DIR=%s", fieldName, c.ErunDir),
	}
}

func ensureString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%s", v)
}

func (c *Config) UnmarshalTOML(data interface{}) error {
	ht, ok := data.(map[string]interface{})
	if !ok {
		return errors.Errorf("Expected a map in Erun config, but got %T", data)
	}
	profile := ensureString(ht["profile"])
	cmd := ensureString(ht["cmd"])
	cfgFile := ensureString(ht["configFile"])
	dirSuffix := ensureString(ht["dirSuffix"])
	erunDir := ensureString(ht["erunDir"])
	err := c.UnmarshalAny(cmd, profile, cfgFile, dirSuffix, erunDir)
	if err != nil {
		err = errors.Errorf("%s when parsing TOML", err)
	}
	return err
}

func (c *Config) SetValue(prefix string, ctx interface{}) error {
	var (
		cmd       string
		profile   string
		cfgFile   string
		dirSuffix string
		erunDir   string
		source    string
	)
	if ctx == nil {
		cmd = os.Getenv(prefix + "_CMD")
		profile = os.Getenv(prefix + "_PROFILE")
		cfgFile = os.Getenv(prefix + "_CONFIG_FILE")
		dirSuffix = os.Getenv(prefix + "_DIR_SUFFIX")
		erunDir = os.Getenv(prefix + "_ERUN_DIR")
		source = "environment"
	} else {
		options := ctx.(map[string]string)
		cmd = options[prefix+".cmd"]
		profile = options[prefix+".profile"]
		cfgFile = options[prefix+".config_file"]
		dirSuffix = options[prefix+".dir_suffix"]
		erunDir = options[prefix+".erun_dir"]
		source = "cli options"
	}
	err := c.UnmarshalAny(cmd, profile, cfgFile, dirSuffix, erunDir)
	if err != nil {
		err = errors.Errorf("%s when parsing %s", err, source)
	}
	return err
}

func (c *Config) UnmarshalAny(cmd, profile, cfgFile, dirSuffix, erunDir string) error {
	if cmd == "" && profile == "" {
		if c == nil {
			*c = Config{}
		}
		if c.Profile == nil {
			c.Profile = &ProfileIO{}
		}
		return nil
	}

	if profile == "" && cmd != "" {
		return errors.Errorf("Don't know what Erun profile to use. " +
			"(Either specify profile, or remove the command")
	}

	if c == nil {
		*c = Config{}
	}

	for _, p := range profiles() {
		if p.Name() == profile {
			c.Profile = p
			c.ConfigFile = cfgFile
			c.DirSuffix = dirSuffix
			c.ErunDir = erunDir
			raw, err := cmdToJSON(cmd, p)
			if err != nil {
				return err
			}
			return c.Profile.UnmarshalJSON(raw)
		}
	}
	if profile == "" {
		c.Profile = &ProfileIO{}
		c.ConfigFile = cfgFile
		c.DirSuffix = dirSuffix
		c.ErunDir = erunDir
		return nil
	}
	return errors.Errorf("Illegal Erun profile %s", profile)
}

// A version of Config that uses only primitive data types and is
// therefore suitable for (un)marshaling.
type primitiveConfig struct {
	DirSuffix   string                 `json:"dir-suffix,omitempty"`
	ErunDir     string                 `json:"erun-dir,omitempty"`
	Profile     map[string]interface{} `json:"profile,omitempty"`
	ProfileName string                 `json:"profile-name"`
	ConfigFile  string                 `json:"config-file,omitempty"`
}

// Yes, it's marshal by value and unmarshal by reference.  Because we
// like Go very much.
func (c Config) MarshalJSON() ([]byte, error) {
	dummy := map[string]interface{}{}
	if c.DirSuffix != "" {
		dummy["dir-suffix"] = c.DirSuffix
	}
	if c.ErunDir != "" {
		dummy["erun-dir"] = c.ErunDir
	}
	if c.ConfigFile != "" {
		dummy["config-file"] = c.ConfigFile
	}
	dummy["profile-name"] = c.Profile.Name()
	dummy["profile"] = c.Profile
	return json.Marshal(dummy)
}

func (c *Config) UnmarshalJSON(data []byte) error {
	var pc primitiveConfig
	if err := json.Unmarshal(data, &pc); err != nil {
		return err
	}

	c.DirSuffix = pc.DirSuffix
	c.ErunDir = pc.ErunDir
	c.ConfigFile = pc.ConfigFile

	for _, p := range profiles() {
		if pc.ProfileName == p.Name() {
			profileData, err := json.Marshal(pc.Profile)
			if err != nil {
				return err
			}
			c.Profile = p
			return json.Unmarshal(profileData, c.Profile)
		}
	}

	return errors.Errorf("Unknown profile kind: %s\n", pc.ProfileName)
}

// Note: We assume that zero values (e.g. 0, false) mean that the user did not supply the values for those fields,
// and therefore we don't pass them to erun. If this assumption is not good enough for a specific field,
// it should be moved to a pointer type such as optional.Int.

////////////////////////////////////////////////////////////

type ProfileIO struct {
	Duration           time.Duration `erun:"duration" default:"180"`
	Clients            int           `erun:"clients"`
	NrFiles            int           `erun:"nr-files"`
	MaxFileSize        size.Size     `erun:"max-file-size" default:"104857600"` // 100 MiB
	QueueSize          int           `erun:"queue-size" default:"16"`
	Readwrites         string        `erun:"readwrites" default:"70"`
	Sequential         bool          `erun:"sequential"`
	DataPayload        bool          `erun:"data-payload"`
	InitialWritePhase  bool          `erun:"initial-write-phase"`
	InitialWriteStop   bool          `erun:"initial-write-stop"`
	ReuseExistingFiles bool          `erun:"reuse-existing-files"`
	MaxIoSize          size.Size     `erun:"max-io-size"`
	MinIoSize          size.Size     `erun:"min-io-size"`
	RecoveryTimeout    time.Duration `erun:"recovery-timeout" default:"20"`
	ShuffleFiles       time.Duration `erun:"shuffle-files-secs"`
	CncTraces          size.Size     `erun:"cnc-traces"`
	MinUncomp          int           `erun:"min-uncomp"`
	MaxUncomp          int           `erun:"max-uncomp"`
	Dedup              int           `erun:"dedup"`
	ValidationPathName string        `erun:"validation-pathname"`
	InitialWiteOffset  size.Size     `erun:"initial-write-offset"`
	ShuffleFilesSecs   int           `erun:"shuffle-files-secs"`
	IoRateLimit        int           `erun:"io-rate-limit"`
	// ReadBeforeWrite    bool          `erun:"read-before-write" default:"true"`
}

func (pio *ProfileIO) Name() string  { return "io" }
func (pio *ProfileIO) ToCmd() string { return toCmd(pio) }
func (pio *ProfileIO) ToFullCmd(opts *ErunFormatCmdOpts) string {
	return opts.Combine(
		pio,
		"--erun-dir", opts.ErunDir(),
		"--json-output", opts.InBucket(opts.ResultsFullLog),
		pio.ToCmd(),
		"; rc=$?",
		opts.LogSummary(
			"; tail -1 %s > %s",
			opts.ResultsFullLog,
			opts.ResultFile,
		),
		opts.CoreDumpsScript(),
		"; exit $rc",
	)
}
func (pio *ProfileIO) HasResults() bool { return true }
func (pio *ProfileIO) UnmarshalJSON(data []byte) error {
	return unmarshalProfile(reflect.ValueOf(pio).Elem(), data)
}

func (pio *ProfileIO) String() string {
	return fmt.Sprintf(
		`{Clients: %v, MaxFilesSize: %v, NrFiles: %v, Readwrites: %v, QueueSize: %v, MaxIoSize: %v, MinIoSize: %v, RecoveryTimeout: %v, InitialWritePhase: %v, InitialWriteStop: %v, ReuseExistingFiles: %v}`,
		pio.Clients, pio.MaxFileSize, pio.NrFiles, pio.Readwrites, pio.QueueSize, pio.MaxIoSize, pio.MinIoSize, pio.RecoveryTimeout, pio.InitialWritePhase, pio.InitialWriteStop, pio.ReuseExistingFiles,
	)
}

////////////////////////////////////////////////////////////

type ProfileMetafuzzer struct {
	Duration         time.Duration `erun:"duration"`
	Clients          int           `erun:"clients"`
	QueueSize        int           `erun:"queue-size"`
	TargetNrDirs     int           `erun:"target-nr-dirs"`
	TargetNrFiles    int           `erun:"target-nr-files"`
	LogFile          string        `erun:"log-file"`
	MaxLogFileSize   size.Size     `erun:"max-log-file-size"`
	EndStateDumpFile string        `erun:"end-state-dump-file"`
	// --[no-]stdout                [BOOL]    [ default:      no ] # emit detailed log to standard output too
	// --[no-]data-verbose          [BOOL]    [ default:     yes ] # dump written data to log
	// --create-dir                 [INTEGER] [ default:    5000 ] #
	// --create-file---default      [INTEGER] [ default:   20000 ] #
	// --destroy-dir                [INTEGER] [ default:    5000 ] #
	// --destroy-file---default     [INTEGER] [ default:   20000 ] #
	// --read-file                  [INTEGER] [ default:  400000 ] #
	// --rename--colliding-dest     [INTEGER] [ default:   10000 ] #
	// --validate--dir-content      [INTEGER] [ default:  220000 ] #
	// --write-file                 [INTEGER] [ default:  320000 ] #
}

func (pmf *ProfileMetafuzzer) Name() string  { return "metafuzzer" }
func (pmf *ProfileMetafuzzer) ToCmd() string { return toCmd(pmf) }
func (pmf *ProfileMetafuzzer) ToFullCmd(opts *ErunFormatCmdOpts) string {
	return opts.Combine(
		pmf,
		"--erun-dir", opts.ErunDir(),
		pmf.ToCmd(),
		opts.CoreDumpsScript(),
		opts.LogSummary(""),
	)
}
func (pmf *ProfileMetafuzzer) HasResults() bool { return false }
func (pmf *ProfileMetafuzzer) UnmarshalJSON(data []byte) error {
	return unmarshalProfile(reflect.ValueOf(pmf).Elem(), data)
}

////////////////////////////////////////////////////////////

type ProfileMetabench struct {
	Description           string `default:"Tesla Metabench test"`
	Scenario              string // e.g. "one"
	QueueSizePerClient    int    // e.g. 2
	NrClients             int    // e.g. 2
	NrDirs                int    // e.g. 4
	NrSubdirs             int    // e.g. 4
	NrTotalFiles          int    // e.g. 20
	NrMoves               int    // e.g. 100
	NrIterationsPerWorker int    // e.g. 2000
}

func (pmb *ProfileMetabench) Name() string  { return "metabench" }
func (pmb *ProfileMetabench) ToCmd() string { return toCmd(pmb) }
func (pmb *ProfileMetabench) ToFullCmd(opts *ErunFormatCmdOpts) string {
	return opts.Combine(
		pmb,
		"--param-file", opts.ParamFile(),
		"--report-file", opts.ReportFile(),
		opts.CoreDumpsScript(),
		opts.LogSummary(""),
	)

}
func (pmb *ProfileMetabench) HasResults() bool { return false }
func (pmb *ProfileMetabench) UnmarshalJSON(data []byte) error {
	return unmarshalProfile(reflect.ValueOf(pmb).Elem(), data)
}

////////////////////////////////////////////////////////////

type ProfileDevRegression struct{}

func (pdr *ProfileDevRegression) Name() string  { return "dev_regression" }
func (pdr *ProfileDevRegression) ToCmd() string { return toCmd(pdr) }
func (pdr *ProfileDevRegression) ToFullCmd(opts *ErunFormatCmdOpts) string {
	return opts.Combine(
		pdr,
		pdr.ToCmd(),
		opts.CoreDumpsScript(),
		opts.LogSummary(""),
	)
}
func (pdr *ProfileDevRegression) HasResults() bool { return false }
func (pdr *ProfileDevRegression) UnmarshalJSON(data []byte) error {
	return unmarshalProfile(reflect.ValueOf(pdr).Elem(), data)
}

////////////////////////////////////////////////////////////////////////////////

// Convert options to an erun command line
func toCmd(p Profile) string {
	if p == nil {
		return ""
	}

	var opts []string

	v := reflect.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag
		name := tag.Get("erun")
		value := v.Field(i).Interface()

		var optionValue string
		// Check if the field's value is zero
		if value == reflect.Zero(v.Field(i).Type()).Interface() {
			// No value was specified. Use the default if there is one
			optionValue = tag.Get("default")
			if optionValue == "" {
				// No default either; skip this field
				continue
			}
		} else {
			optionValue = valueToOption(value)
		}

		// Append the option name
		opts = append(opts, fmt.Sprintf("--%v", name))
		if optionValue != "" {
			opts = append(opts, optionValue)
		}
	}

	return strings.Join(opts, " ")
}

func valueToOption(value interface{}) string {
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		return "" // Just specify the option without any value
	case string:
		return fmt.Sprintf("%q", v)
	case time.Duration:
		if v < 1*time.Second {
			panic(fmt.Sprintf("Unsupported value for duration (< 1 second): %v", v))
		}
		return fmt.Sprintf("%d", int(v.Seconds()))
	case size.Size:
		return fmt.Sprintf("%d", v/size.Bytes) // TODO: Convert to user-readable units such as "100M" for convenience?
	default:
		panic(fmt.Sprintf("Unsupported type '%T' for value '%[1]v'", v))
	}
}

func parseOptions(options string) (map[string]string, []string) {
	parsed := map[string]string{}
	var (
		key, value string
		args       []string
	)

	// adding " -" to earn another iteration which will cause updating last pair
	for _, match := range tokenize(options + " -") {
		if strings.HasPrefix(match, "-") {
			if key != "" {
				parsed[key] = value
				value = ""
			}
			key = match
		} else if value == "" {
			value = match
		} else {
			args = append(args, match)
		}
	}
	return parsed, args
}

func tokenize(options string) []string {
	start, end := 0, 0
	quot, apos, escape, started := false, false, false, false
	var (
		c      rune
		result []string
	)

	collect := func(pos int) {
		result = append(result, options[start:pos])
		started = false
	}
	begin := func() {
		if !started {
			start = end
		}
		started = true
	}

	for end, c = range options {
		switch c {
		case '"':
			if !(escape || apos) {
				if quot {
					quot = false
					collect(end + 1)
				} else {
					quot = true
					begin()
				}
			}
			escape = false
		case '\'':
			if !(escape || quot) {
				if apos {
					apos = false
					collect(end + 1)
				} else {
					apos = true
					begin()
				}
			}
			escape = false
		case '\\':
			escape = !escape
			begin()
		case ' ', '\t', '\r', '\n':
			if !(quot || apos || escape) && started {
				collect(end)
				started = false
			}
			escape = false
		default:
			begin()
		}
	}
	collect(end + 1)
	return result
}

// Removing the following flags passed by old jenkins jobs:
// 'profile' - should be defined by the conf param 'TESLA_ERUN_PROFILE'
// 'erunDir' - we handle it by ourselves
// 'log-file'- Not saving detailed log file, we log the stdout
//
func cmdToJSON(cmdflags interface{}, p Profile) ([]byte, error) {
	var cmd string
	if scmd, ok := cmdflags.(string); ok {
		cmd = scmd
	} else {
		cmd = fmt.Sprintf("%s", cmdflags)
	}
	options, args := parseOptions(cmd)
	if args != nil {
		fmt.Printf("Extra elements after parsing Erun options: %s", args)
	}
	for k, v := range options {
		if k == "{}" || v == "{}" {
			// TODO(olegs): Print a warning about not using placeholders
			delete(options, k)
			continue
		}
		switch k {
		case "-p", "--profile", "--erun-dir":
			// TODO(olegs): Print a warning about unsupported options
			delete(options, k)
		case "--log-file":
			// TODO(olegs): Print a warning about automatic replacements
			options[k] = erunLogsOnHost
		}
	}

	return json.Marshal(camelCase(options, p))
}

func camelCase(options map[string]string, p Profile) map[string]interface{} {
	result := map[string]interface{}{}

	for k, v := range options {
		ck := toCamelCase(k)
		result[ck] = decode(ck, v, p)
	}
	return result
}

func decode(fieldName, encoded string, p Profile) interface{} {
	durType := reflect.ValueOf(time.Second).Type()
	switch encoded {
	case "true":
		return true
	case "false":
		return false
	case "null":
		return nil
	case "":
		return true
	}
	if i, err := strconv.Atoi(encoded); err == nil {
		if reflect.Indirect(reflect.ValueOf(p)).FieldByName(fieldName).Type() == durType {
			// TODO(olegs): Are all times given in seconds?  Why do we need
			// these fields to be in time.Duration?
			return time.Duration(i) * time.Second
		}
		return i
	}
	if s, err := size.Parse(encoded); err == nil {
		return s
	}
	return encoded
}

func toCamelCase(option string) string {
	result := make([]rune, len(option))
	nextCap := true
	j := 0
	for _, c := range option {
		if c == '-' {
			nextCap = true
		} else {
			if nextCap {
				result[j] = unicode.ToUpper(c)
				nextCap = false
			} else {
				result[j] = c
			}
			j++
		}
	}
	return string(result[:j])
}

func unmarshalProfile(profile reflect.Value, data []byte) error {
	dummy := map[string]interface{}{}
	if err := json.Unmarshal(data, &dummy); err != nil {
		return err
	}
	sizeType := reflect.ValueOf(size.Byte).Type().Name()
	durationType := reflect.ValueOf(time.Second).Type().Name()
	intType := reflect.ValueOf(int(0)).Type().Name()
	stringType := reflect.ValueOf("str").Type().Name()

	for k, v := range dummy {
		field := profile.FieldByName(k)
		if !field.IsValid() {
			return errors.Errorf(
				"Invalid Erun option: %s for profile %s",
				k, profile.Type(),
			)
		}
		switch field.Type().Name() {
		case intType:
			if f, ok := v.(float64); ok {
				field.Set(reflect.ValueOf(int(f)))
			} else {
				field.Set(reflect.ValueOf(v))
			}
		case sizeType:
			// Yey Go!
			switch av := v.(type) {
			case int64:
				field.Set(reflect.ValueOf(size.Size(av)))
			case int32:
				field.Set(reflect.ValueOf(size.Size(av)))
			case uint32:
				field.Set(reflect.ValueOf(size.Size(av)))
			case uint64:
				field.Set(reflect.ValueOf(size.Size(av)))
			case int:
				field.Set(reflect.ValueOf(size.Size(av)))
			case float64:
				field.Set(reflect.ValueOf(size.Size(av)))
			case float32:
				field.Set(reflect.ValueOf(size.Size(av)))
			case string:
				sz, err := size.Parse(av)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(sz))
			default:
				return errors.Errorf("Cannot parse %v into %s", v, sizeType)
			}
		case durationType:
			switch av := v.(type) {
			case int64:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case int32:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case uint32:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case uint64:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case int:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case float64:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case float32:
				field.Set(reflect.ValueOf(time.Duration(av)))
			case string:
				d, err := time.ParseDuration(av)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(d))
			default:
				return errors.Errorf("Cannot parse %v into %s", v, durationType)
			}
		case stringType:
			field.Set(reflect.ValueOf(fmt.Sprintf("%v", v)))
		default:
			field.Set(reflect.ValueOf(v))
		}
	}
	return nil
}
