package portmapper

import "fmt"

func (m Mapping) String() string {
	return fmt.Sprintf("Mapping: Program %d, Version %d, Protocol %d, Port %d",
		m.Prog, m.Vers, m.Prot, m.Port)
}
