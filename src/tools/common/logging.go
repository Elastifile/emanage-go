package common

import (
	"time"

	"jobs"
	"golang.org/x/crypto/ssh/agent"

	filestore_api "github.com/elastifile/emanage-go/src/filestore/api"
	"github.com/elastifile/emanage-go/src/types"
)

type DefaultLoggingTool struct {
	*DefaultExternalTool
	Props types.LoggingProperties
}

func NewLoggingTool(props types.LoggingProperties) *DefaultLoggingTool {
	return &DefaultLoggingTool{
		DefaultExternalTool: NewExternalTool(props),
		Props:               props,
	}
}

func (t *DefaultLoggingTool) Stop(context *types.Context) error {
	logger.Info("logging.Stop", "tool", t.Props.Name())
	job, err := jobs.ByID(t.Props.Params().JobID)
	if err != nil {
		return err
	}
	err = t.DefaultExternalTool.Stop(context)
	logger.Info("Will validate results", "tool", t.Props.Name(), "error", err)
	verr := t.ValidateResults(&job)
	if err != nil {
		return err
	}
	return verr
}

func (t *DefaultLoggingTool) ValidateResults(job *types.Job) (err error) {
	verbose := false
	msg := agent.Msg()
	conf := agent.Config()

	timeout := time.Minute
	for start := time.Now(); time.Since(start) < timeout; time.Sleep(5 * time.Second) {
		fs, e := filestore_api.NewClient(msg, conf.FilestoreHostInternal(), verbose)
		if e != nil {
			continue
		}

		filesByHost, e := fs.RequestRetrieveFiles(
			t.DataStorePath(),
			t.Props.GetResultFilesPatterns(),
			conf.AllLoadersInternal(),
		)
		if e != nil {
			continue
		}

		// Handle results
		err = t.Props.GetResults(&types.ResultOpts{
			FilesByHost: filesByHost,
			Conf:        job.Config,
			JobID:       job.ID,
		})

		if err == nil {
			break
		}
	}

	if err != nil {
		_ = jobs.MarkFail(job.ID, err)
	}

	return err
}
