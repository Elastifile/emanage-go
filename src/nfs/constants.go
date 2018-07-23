package nfs

// Mask off just the documented mode bits.
// Linux knfsd sometimes sets the higher bits as well (e.g. 0100666 instead of 0666).
// (this has been fixed in newer Linux kernels, see the history of linux/fs/nfsd/nfs3xdr.c)
const ModeMask = 0x00000fff
