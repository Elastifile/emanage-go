package portmapper

import "fmt"

type ProgramNotAvailableError struct {
	Mapping
}

func (e *ProgramNotAvailableError) Error() string {
	return fmt.Sprintf("Program not available: %v", e.Mapping)
}
