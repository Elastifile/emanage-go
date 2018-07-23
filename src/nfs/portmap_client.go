package nfs

import (
	"nfs/sunrpc/client"
	"nfs/sunrpc/portmapper"
)

// PortmapClient represents a client for the PORTMAP protocol.
type PortmapClient struct {
	Pmap2 *portmapper.PmapV2
}

// NewPortmapClient creates a new PortmapClient.
func NewPortmapClient(conf client.RPCClientConfig) (*PortmapClient, error) {
	conf.Name = "portmap"
	conf.NumConns = 1
	conf.Port = portmapper.PMAP_PORT
	c, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &PortmapClient{
		Pmap2: portmapper.NewPmapV2(c),
	}, nil
}

// GetPort looks up an returns the port number of a specific SunRPC program.
func (p *PortmapClient) GetPort(program uint32) (port uint32, err error) {
	mapping := portmapper.Mapping{
		Prog: program,
		Vers: 3,
		Prot: portmapper.IPPROTO_TCP,
		Port: 0,
	}

	pPort, err := p.Pmap2.Getport(&mapping)
	if err != nil {
		return 0, err
	}

	if *pPort == 0 {
		return 0, &portmapper.ProgramNotAvailableError{Mapping: mapping}
	}

	return *pPort, nil
}
