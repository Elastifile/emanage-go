package fio

// TODO: When relevant fields will determine placeholders will be added.
const defaultFileTemplate = `
; default jobs: normal,compress,dedup,sequential
[global]
blocksize=4k
rw=randrw
size=64m
time_based=1
direct=1
verify=sha1
runtime=120
bssplit=4k/70:64k/10:8k/10:2k/10

[normal]

[compress]
buffer_compress_percentage=70

[dedup]
dedupe_percentage=70

[sequential]
rw=rw`

// Data integrity template
const diFileTemplate = `
; data integrity jobs: write-phase,verify-phase
; writes 512 byte verification blocks until the disk is full,
; then verifies written data
[global]
thread=1
bs=64k
direct=1
ioengine=sync
verify=sha256
verify_interval=512
verify_fatal=1
size=100m
runtime=1800

[write-phase]
; filename=datafile.tmp	; or use a full disk, for example /dev/sda
rw=write
; fill_device=1
do_verify=0
nrfiles=100
filename_format=datafile.tmp.$filenum

[verify-phase]
stonewall
create_serialize=0
; filename=datafile.tmp
rw=read
do_verify=1
nrfiles=100
filename_format=datafile.tmp.$filenum`
