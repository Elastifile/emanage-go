package common

import (
	"path/filepath"

	"github.com/go-errors/errors"

	"types"
)

const DataVolume = "/data/"

var ErrDoNotRun = errors.New("Do not run the tool on this host")

// UnifiedResultsHeader used as the header for generated csv results
var UnifiedResultsHeader = []string{"Identifier", "ToolName", "Workload", "RequestedIOPs", "IOPs", "Latency(ms)", "Bandwidth", "JobID"}

func PkgDataStorePath(jobID types.JobID, PkgName string) string {
	return filepath.Join(jobDataStorePath(jobID), PkgName)
}

func jobDataStorePath(jobID types.JobID) string {
	return filepath.Join(DataVolume, string(jobID))
}

func ValidateArgs(params *types.ToolParams) (err error) {
	if len(params.Args) < 1 {
		err = errors.New("no args supplied, image should be first arg")
	}
	return err
}

func ValidateExports(params *types.ToolParams) (err error) {
	conf := params.Config
	if conf.Tesla.Elfs.Export == "" {
		err = errors.Errorf("No ELFS export was specified, conf: %+v", conf)
	}
	return err
}
