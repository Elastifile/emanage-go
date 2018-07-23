package sfs2008

const template = `LOAD="{{.Sfs2008.LoadIo}}"
INCR_LOAD={{.Sfs2008.IncrLoad}}
NUM_RUNS={{.Sfs2008.NumberOfRuns}}
PROCS={{.Sfs2008.ProcessesPerClient}}
CLIENTS="{{join .LoadersStr " "}}"
MNT_POINTS="{{$mountpoint := print .System.Frontend ":" .Tesla.Elfs.Export}}{{$delimiter:=" "}}{{repeatStr $mountpoint .Sfs2008.ProcessesPerClient $delimiter}}"
BIOD_MAX_WRITES=2
BIOD_MAX_READS=2
IPV6_ENABLE="off"
FS_PROTOCOL="nfs"
SFS_DIR="src"
SUFFIX=""
WORK_DIR="result"
PRIME_MON_SCRIPT=""
PRIME_MON_ARGS=""
INIT_TIMEOUT=8000
BLOCK_SIZE=32
SFS_NFS_USER_ID=500
SFS_NFS_GROUP_ID=500
RUNTIME={{.Sfs2008.Runtime}}
WARMUP_TIME={{.Sfs2008.WarmupTime}}
ACCESS_PCNT=30
APPEND_PCNT=70
BLOCK_FILE=""
DIR_COUNT=30
FILE_COUNT=
SYMLINK_COUNT=20
TCP="on"
DEBUG=""
DUMP=
POPULATE=
LAT_GRAPH=
PRIME_SLEEP=0
PRIME_TIMEOUT=0
{{with .Sfs2008.MixFile}}MIXFILE="{{.}}"{{end}}
`
