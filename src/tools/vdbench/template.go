package vdbench

var vdBenchTemplate = `
# Case will define hd's and fsd's according to the way this test is run
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
fsd={{.Name}},anchor={{.Anchor}},depth=2,width=4,files=10,size=(128k,30,512k,30,1m,25,10m,10,20m,5),shared=yes,openflags=o_direct
{{end}}

###  256k Sequential IO
fwd=256K_Sequential_Write,fsd=fsd*,operation=write,xfersize=256k,fileio=sequential,fileselect=random,threads=8

# Run Definitions
rd=4K_8K_32K_Random_R70_W30,fwd=256K_Sequential_Write,foroperations=(write,read),fwdrate=max,format=restart,elapsed=60,interval=1,pause=10
`
