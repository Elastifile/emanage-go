package nfsx

import "fmt"

func (fh Fh3) String() string {
	return fmt.Sprintf("0x%x", fh.Data)
}
