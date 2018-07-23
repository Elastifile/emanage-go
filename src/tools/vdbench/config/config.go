package vdbench_config

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"size"
)

type Config struct {
	ConfigFile string `yaml:"ConfigFile" tesla:"filename"`
	Case       Case   `yaml:"Case"`
}

type Host struct {
	User   string `yaml:"User"`
	Shell  string `yaml:"Shell"`
	Name   string `yaml:"Name"`
	System string `yaml:"System"`
}

func (h *Host) String() string {
	return fmt.Sprintf("hd=%s,user=%s,shell=%s,system=%s",
		h.Name, h.User, h.Shell, h.System)
}

type FileSystem struct {
	Name      string `yaml:"Name"`
	Anchor    string `yaml:"Anchor"`
	Depth     int    `yaml:"Depth"`
	Width     int    `yaml:"Width"`
	Files     int    `yaml:"Files"`
	Size      []int  `yaml:"Size"`
	Shared    bool   `yaml:"Shared"`
	OpenFlags string `yaml:"OpenFlags"`
}

func stringSizes(sizes []int) string {
	printed := make([]string, len(sizes))
	for i, o := range sizes {
		s := size.Size(o)
		printed[i] = s.String()
	}
	return strings.Join(printed, ",")
}

func stringBool(val bool) string {
	if val {
		return "yes"
	} else {
		return "no"
	}
}

func (fs *FileSystem) String() string {
	return fmt.Sprintf("fsd=%s,anchor=%s,depth=%d,width=%d,files=%d,size=(%s),shared=%s,openflags=%s",
		fs.Name, fs.Anchor, fs.Depth, fs.Width, fs.Files,
		stringSizes(fs.Size), stringBool(fs.Shared), fs.OpenFlags)
}

type Workload struct {
	Name        string `yaml:"Name"`
	FileSystems string `yaml:"FileSystems"`
	Operation   string `yaml:"Operation"`
	Xfersize    int    `yaml:"Xfersize"`
	FileIO      string `yaml:"FileIO"`
	FileSelect  string `yaml:"FileSelect"`
	Threads     int    `yaml:"Threads"`
}

func (w *Workload) String() string {
	return fmt.Sprintf("fwd=%s,fsd=%s,operation=%s,xfersize=%s,fileio=%s,fileselect=%s,threads=%d",
		w.Name, w.FileSystems, w.Operation,
		size.Size(w.Xfersize).String(),
		w.FileIO, w.FileSelect, w.Threads)
}

type Run struct {
	Name          string   `yaml:"Name"`
	Workload      string   `yaml:"Workload"`
	ForOperations []string `yaml:"ForOperations"`
	WorkloadRate  string   `yaml:"WorkloadRate"`
	Format        string   `yaml:"Format"`
	Elapsed       int      `yaml:"Elapsed"`
	Interval      int      `yaml:"Interval"`
	Pause         int      `yaml:"Pause"`
}

func (r *Run) String() string {
	return fmt.Sprintf("rd=%s,fwd=%s,foroperations=(%s),fwdrate=%s,format=%s,elapsed=%d,interval=%d,pause=%d",
		r.Name, r.Workload, strings.Join(r.ForOperations, ","),
		r.WorkloadRate, r.Format, r.Elapsed, r.Interval, r.Pause)
}

type Case struct {
	DedupRatio  int          `default:"10" yaml:"DedupRatio"`
	DedupUnit   int          `default:"2048" yaml:"DedupUnit"`
	CompRatio   float64      `default:"4.0" yaml:"CompRatio"`
	Debug       int          `default:"25" yaml:"Debug"`
	Hosts       []Host       `yaml:"Hosts"`
	FileSystems []FileSystem `yaml:"FileSystems"`
	Workloads   []Workload   `yaml:"Workloads"`
	Runs        []Run        `yaml:"Runs"`
}

func (r *Case) String() string {
	tpl := template.New("cmd")
	funcMap := template.FuncMap{
		"size": func(x int) string {
			return size.Size(x).String()
		},
	}

	parse, _ := tpl.Funcs(funcMap).Parse(`dedupratio={{.DedupRatio}}
dedupunit={{size .DedupUnit}}
compratio={{.CompRatio}}
debug={{.Debug}}

## Hosts definitions
{{range .Hosts}}{{.}}
{{end}}

## Filesystem definitions
{{range .FileSystems}}{{.}}
{{end}}

## Workloads definitions
{{range .Workloads}}{{.}}
{{end}}

## Runs definitions
{{range .Runs}}{{.}}
{{end}}`)
	var buf bytes.Buffer
	_ = parse.Execute(&buf, r)
	return buf.String()
}

func New(testDef string) *Config {
	return &Config{
		ConfigFile: testDef,
	}
}

func NewCase() *Case {
	return &Case{
		DedupRatio: 8,
		DedupUnit:  int(2 * size.KiB),
		CompRatio:  10,
		Debug:      25,
	}
}
