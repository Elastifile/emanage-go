package portmapper

import (
	"bytes"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	xdr "github.com/davecgh/go-xdr/xdr2"
)

func TestPmaplistFirstBasic(*testing.T) {
	list := &PmaplistFirst{
		Next: &Pmaplist{
			Map: Mapping{
				Prog: 10,
				Vers: 20,
				Prot: 30,
				Port: 40,
			},
			Next: &Pmaplist{
				Map: Mapping{
					Prog: 11,
					Vers: 12,
					Prot: 13,
					Port: 14,
				},
				Next: &Pmaplist{
					Map: Mapping{
						Prog: 21,
						Vers: 22,
						Prot: 23,
						Port: 24,
					},
				},
			},
		},
	}
	testPmaplistFirst(list)
}

func TestPmaplistFirstWrapped(*testing.T) {
	mappingsIn := []Mapping{
		{Prog: 110, Vers: 20, Prot: 30, Port: 40},
		{Prog: 111, Vers: 12, Prot: 13, Port: 14},
		{Prog: 121, Vers: 22, Prot: 23, Port: 24},
	}

	// Convert from slice to list
	var listIn *PmaplistFirst

	listIn = &PmaplistFirst{}

	current := &listIn.Next
	for _, m := range mappingsIn {
		*current = &Pmaplist{Map: m}
		current = &(*current).Next
	}

	listOut := testPmaplistFirst(listIn)
	spew.Dump(listOut)

	// Convert back from list to slice
	mappingsOut := []Mapping{}
	currentOut := listOut.Next
	for currentOut != nil {
		mappingsOut = append(mappingsOut, currentOut.Map)
		currentOut = currentOut.Next
	}

	log.Printf("mappingsOut: %+v", mappingsOut)
}

func testPmaplistFirst(in *PmaplistFirst) (out *PmaplistFirst) {
	var buf bytes.Buffer
	var err error

	encoder := xdr.NewEncoder(&buf)
	decoder := xdr.NewDecoder(&buf)

	out = &PmaplistFirst{}

	log.Printf("encode")
	log.Printf("in: %+v", in)
	err = in.Encode(encoder)
	log.Printf("err: %v", err)

	log.Printf("buf: %v", buf.Bytes())

	log.Printf("decode")
	err = out.Decode(decoder)
	log.Printf("err: %v", err)
	log.Printf("out: %+v", out)

	return out
}
