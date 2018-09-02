package nfs

import (
	log "gopkg.in/inconshreveable/log15.v2"
	logext "gopkg.in/inconshreveable/log15.v2/ext"

	"github.com/elastifile/emanage-go/src/nfs/sunrpc/client"
	"github.com/elastifile/emanage-go/src/nfs/sunrpc/nfsx"
	"github.com/elastifile/emanage-go/src/nfs/sunrpc/portmapper"
	"github.com/elastifile/emanage-go/src/nfs/sunrpc/rpc2"
	"github.com/elastifile/emanage-go/src/types"
)

// NfsClient represents an NFS client.
type NfsClient struct {
	Auth   rpc2.Auth
	Pmap2  *portmapper.PmapV2
	Mount3 *nfsx.MountprogmountV3
	Nfs3   *nfsx.NfsV3
	*Limits

	host     types.Host
	pmap     *PortmapClient
	numConns int

	log.Logger
}

// NewNfsClient creates a new NfsClient, including SunRPC clients for the portmap, mount and NFS protocols.
func NewNfsClient(conf client.RPCClientConfig) (*NfsClient, error) {
	return NewNfsClientFromExisting(conf, nil)
}

func NewNfsClientFromExisting(conf client.RPCClientConfig, existingClient *NfsClient) (*NfsClient, error) {
	var (
		pmap  *PortmapClient
		mount *client.Client
		cNfs  *client.Client
	)

	if existingClient != nil {
		pmap = existingClient.pmap
		if existingClient.Mount3 != nil {
			mount = existingClient.Mount3.Client
		}
		if existingClient.Nfs3 != nil {
			cNfs = existingClient.Nfs3.Client
		}
	}

	var err error

	if pmap == nil {
		pmapCfg := conf
		pmap, err = NewPortmapClient(pmapCfg)
		if err != nil {
			return nil, err
		}
		logger.Debug("Created Portmap", "client", *pmap)
	}

	if mount == nil {
		if conf.Mountport == 0 {
			conf.Mountport, err = pmap.GetPort(nfsx.ProgramMountprogmountV3)
			if err != nil {
				return nil, err
			}
		}
		mountCfg := conf
		mountCfg.Port = conf.Mountport
		mountCfg.NumConns = 1
		mount, err = client.NewClient(mountCfg)
		if err != nil {
			return nil, err
		}
		logger.Debug("Created mount", "client", *mount)
	}

	if cNfs == nil {
		if conf.Port == 0 {
			conf.Port, err = pmap.GetPort(nfsx.ProgramNfsV3)
			if err != nil {
				return nil, err
			}
		}

		nfsCfg := conf
		cNfs, err = client.NewClient(nfsCfg)
		if err != nil {
			return nil, err
		}
		logger.Debug("Created cNfs", "client", *cNfs)
	}

	nc := &NfsClient{
		Auth:     conf.Auth,
		Pmap2:    pmap.Pmap2,
		Mount3:   nfsx.NewMountprogmountV3(mount),
		Nfs3:     nfsx.NewNfsV3(cNfs),
		Logger:   logger.New("id", logext.RandId(8)),
		host:     types.Host(conf.Host),
		pmap:     pmap,
		numConns: conf.NumConns,
	}
	logger.Debug("New NFS", "client", *nc)

	return nc, nil
}

func (nfs *NfsClient) ConnectionCount() int {
	return nfs.numConns
}

// Close SunRPC clients for all protocols.
func (nfs *NfsClient) Close() {
	nfs.Pmap2.Client.Close()
	nfs.pmap.Pmap2.Client.Close()
	nfs.Nfs3.Client.Close()
	nfs.Mount3.Client.Close()
}

// ExpectServerSideClose prepares for the connections being closed by the NFS
// server (e.g. due to a reboot) so that no errors are generated on the client side.
func (nfs *NfsClient) ExpectServerSideClose() {
	nfs.Pmap2.Client.ExpectServerSideClose()
	nfs.Nfs3.Client.ExpectServerSideClose()
	nfs.Mount3.Client.ExpectServerSideClose()
}

// Mount a specific NFS export and return its root filehandle.
func (nfs *NfsClient) Mount(dir string) (fh *nfsx.Fh3, err error) {
	result, err := nfs.Mount3.Mountproc3Mnt((*nfsx.Dirpath)(&dir))
	if err != nil {
		return fh, err
	}
	if result.FhsStatus != nfsx.MNT3_OK {
		return fh, &nfsx.MountError{Status: result.FhsStatus}
	}

	resOk := result.Union.(nfsx.Mountres3Ok)
	return &resOk.Fhandle, nil
}

// Mount a specific NFS export and return its root filehandle.
func (nfs *NfsClient) Unmount(dir string) error {
	return nfs.Mount3.Mountproc3Umnt((*nfsx.Dirpath)(&dir))
}

// Lookup the NFS filehandle for a specific entry in a directory.
func (nfs *NfsClient) Lookup(dirFh *nfsx.Fh3, name nfsx.Filename3) (fh *nfsx.Fh3, attr *nfsx.Fattr3, err error) {
	result, err := nfs.Nfs3.Proc3Lookup(&nfsx.Lookup3args{
		What: nfsx.Diropargs3{Dir: *dirFh, Name: name},
	})
	if err != nil {
		return
	}
	if result.Status != nfsx.V3_OK {
		err = &nfsx.NfsError{Status: result.Status}
		return
	}

	resOk := result.Union.(nfsx.Lookup3resok)
	fh = &resOk.Object

	if !resOk.ObjAttributes.AttributesFollow {
		return
	}

	attrData := resOk.ObjAttributes.Union.(nfsx.Fattr3)
	attr = &attrData

	return
}
