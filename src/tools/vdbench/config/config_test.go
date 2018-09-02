package vdbench_config

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/elastifile/emanage-go/src/size"
)

func populateConfig() *Case {
	caseConf := &Case{
		DedupRatio: 10,
		DedupUnit:  int(2 * size.KiB),
		CompRatio:  4.0,
		FileSystems: []FileSystem{
			FileSystem{
				Name:   "fsd1",
				Anchor: "/mnt/elfs",
				Depth:  2,
				Width:  8,
				Files:  150,
				Size: []int{
					int(128 * size.KiB),
					int(30 * size.Byte),
					int(512 * size.KiB),
					int(30 * size.KiB),
					int(1 * size.MiB),
					int(25 * size.Byte),
					int(10 * size.MiB),
					int(10 * size.Byte),
					int(20 * size.MiB),
					int(5 * size.Byte),
				},
				Shared:    true,
				OpenFlags: "o_direct",
			},
		},
	}
	loaders := []string{"loader5a", "loader5b", "loader5c", "loader5d"}
	hds := make([]Host, len(loaders))
	for i, loader := range loaders {
		hds[i] = Host{
			User:   "root",
			Shell:  "ssh",
			Name:   loader,
			System: loader,
		}
	}
	caseConf.Hosts = hds
	return caseConf
}

func TestPrinting(t *testing.T) {
	caseConf := &Case{
		DedupRatio: 10,
		DedupUnit:  int(2 * size.KiB),
		CompRatio:  4.0,
		Hosts: []Host{
			Host{
				User:   "root",
				Shell:  "ssh",
				Name:   "loader9a",
				System: "loader9a",
			},
		},
		FileSystems: []FileSystem{
			FileSystem{
				Name:   "fsd1",
				Anchor: "/data1",
				Depth:  2,
				Width:  8,
				Files:  150,
				Size: []int{
					int(128 * size.KiB),
					int(30 * size.Byte),
					int(512 * size.KiB),
					int(30 * size.KiB),
					int(1 * size.MiB),
					int(25 * size.Byte),
					int(10 * size.MiB),
					int(10 * size.Byte),
					int(20 * size.MiB),
					int(5 * size.Byte),
				},
				Shared:    true,
				OpenFlags: "o_direct",
			},
		},
		Workloads: []Workload{
			Workload{
				Name:        "fwd1",
				FileSystems: "fsd*",
				Operation:   "write",
				Xfersize:    int(256 * size.KiB),
				FileIO:      "sequential",
				FileSelect:  "random",
				Threads:     15,
			},
		},
		Runs: []Run{
			Run{
				Name:          "rd1",
				Workload:      "fwd1",
				ForOperations: []string{"write,read"},
				WorkloadRate:  "max",
				Format:        "restart",
				Elapsed:       6000,
				Interval:      1,
				Pause:         10,
			},
		},
	}
	expected := `dedupratio=10
dedupunit=2 KiB
compratio=4
debug=0

## Hosts definitions
hd=loader9a,user=root,shell=ssh,system=loader9a


## Filesystem definitions
fsd=fsd1,anchor=/data1,depth=2,width=8,files=150,size=(128 KiB,30 bytes,512 KiB,30 KiB,1 MiB,25 bytes,10 MiB,10 bytes,20 MiB,5 bytes),shared=yes,openflags=o_direct


## Workloads definitions
fwd=fwd1,fsd=fsd*,operation=write,xfersize=256 KiB,fileio=sequential,fileselect=random,threads=15


## Runs definitions
rd=rd1,fwd=fwd1,foroperations=(write,read),fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
`
	actual := caseConf.String()
	if expected != actual {
		t.Fatalf(`Expected:
%s
----------------
Got:
%s
----------------
`, expected, actual)
	}
}

func TestTemplate(t *testing.T) {
	content := `# Case will define hd's and fsd's according to the way this test is run
# 
dedupratio={{.DedupRatio}}
dedupunit={{.DedupUnit}}
compratio={{.CompRatio}}
debug={{.Debug}}

# hd's are generated based  on the number of loaders available
# 
hd=default,user=root,shell=ssh
{{range .Hosts}}
hd={{.Name}},system={{.Name}}
{{end}}

# All fsd's are given names fsdN, where N is a positive integer
# 
{{range .FileSystems}}
fsd={{.Name}},anchor={{.Anchor}},depth=2,width=8,files=150,size=(128k,30,512k,30,1m,25,10m,10,20m,5),shared=yes,openflags=o_direct
{{end}}

###  256k Sequential IO
fwd=256K_Sequential_Write,fsd=fsd*,operation=write,xfersize=256k,fileio=sequential,fileselect=random,threads=15
fwd=256K_Sequential_Read,fsd=fsd*,operation=read,xfersize=256k,fileio=sequential,fileselect=random,threads=15

###  256k Random IO
fwd=256K_Random_Write,fsd=fsd*,operation=write,xfersize=256k,fileio=random,fileselect=random,threads=15
fwd=256K_Random_Read,fsd=fsd*,operation=read,xfersize=256k,fileio=random,fileselect=random,threads=15


###  128k Sequential IO
fwd=128K_Sequential_Write,fsd=fsd*,operation=write,xfersize=128k,fileio=sequential,fileselect=random,threads=15
fwd=128K_Sequential_Read,fsd=fsd*,operation=read,xfersize=128k,fileio=sequential,fileselect=random,threads=15


###  128k Random IO
fwd=128K_Random_Write,fsd=fsd*,operation=write,xfersize=128k,fileio=random,fileselect=random,threads=15
fwd=128K_Random_Read,fsd=fsd*,operation=read,xfersize=128k,fileio=random,fileselect=random,threads=15


#### 4K, 8K and 32K/ random / 70% read
fwd=4K_8K_32K_Random_R70_W301,fsd=fsd*,xfersize=(4k,60,8K,30,32k,10),rdpct=70,fileio=random,fileselect=random,threads=60

#### 4K IOPS
# 4K Random Read
fwd=4K_Random_Read1,fsd=fsd*,operation=read,xfersize=4k,fileio=(random),fileselect=random,threads=60
#fwd=4K_Random_Read1,fsd=fsd*,operation=read,xfersize=4k,fileio=(random,shared),fileselect=random,threads=60
# 4K Random Write
fwd=4K_Random_Write1,fsd=fsd*,operation=write,xfersize=4k,fileio=random,fileselect=random,threads=60
# 4K Sequential Read
fwd=4K_Sequential_Read1,fsd=fsd*,operation=read,xfersize=4k,fileio=sequential,fileselect=random,threads=60
# 4K Sequential Write
fwd=4K_Sequential_Write1,fsd=fsd*,operation=write,xfersize=4k,fileio=sequential,fileselect=random,threads=60

#### 2K IOPS
# 2K Random Read Hurt
fwd=2K_Random_Read_Hurt1,fsd=fsd*,operation=read,xfersize=2k,fileio=(random),fileselect=random,threads=60


# Run Definitions
rd=4K_8K_32K_Random_R70_W30,fwd=4K_8K_32K_Random_R70_W301,foroperations=(write,read),fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=128K_Sequential_Write,fwd=128K_Sequential_Write,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=128K_Sequential_Read,fwd=128K_Sequential_Read,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=128K_Random_Write,fwd=128K_Random_Write,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=128K_Random_Read,fwd=128K_Random_Read,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10

rd=256K_Sequential_Write,fwd=256K_Sequential_Write,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=256K_Sequential_Read,fwd=256K_Sequential_Read,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=256K_Random_Write,fwd=256K_Random_Write,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=256K_Random_Read,fwd=256K_Random_Read,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10

rd=2K_Random_Read_Hurt,fwd=2K_Random_Read_Hurt1,fwdrate=max,format=restart,elapsed=60,interval=1,pause=10

rd=4K_Random_Read,fwd=4K_Random_Read1,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=4K_Random_Write,fwd=4K_Random_Write1,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=4K_Sequential_Read,fwd=4K_Sequential_Read1,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10
rd=4K_Sequential_Write,fwd=4K_Sequential_Write1,fwdrate=max,format=restart,elapsed=6000,interval=1,pause=10

`
	tpl := template.New("cmd")
	parse, err := tpl.Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	err = parse.Execute(&buf, populateConfig())
	if err != nil {
		t.Fatal(err)
	}
	_, err = fmt.Println(buf.String())
	if err != nil {
		t.Fatal(err)
	}
}
