package nfsx

import "fmt"

type MountError struct {
	Status mountstat3
}

func (e *MountError) Error() string {
	return fmt.Sprintf("MountError: %v", e.Status)
}

type NfsError struct {
	Status stat3
}

func (e *NfsError) Error() string {
	return fmt.Sprintf("NfsError: %v", e.Status)
}

// ComplianceError happens when server behavior is not according to the spec
type ComplianceError struct {
	Message string
}

func (e *ComplianceError) Error() string {
	return fmt.Sprintf("ComplianceError: %v", e.Message)
}

type AttributeMismatchError struct {
	Prefix string
	*Fattr3
	*Sattr3
}

func (e *AttributeMismatchError) Error() string {
	return fmt.Sprintf("AttributeError: %v: Got Fattr3: %v, Expected Sattr3: %v",
		e.Prefix, e.Fattr3, e.Sattr3)
}
