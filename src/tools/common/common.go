package common

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/go-errors/errors"
	log "gopkg.in/inconshreveable/log15.v2"
	"golang.org/x/crypto/ssh/agent"

	"textutil"
	"github.com/elastifile/emanage-go/src/filestore"
	filestore_types "github.com/elastifile/emanage-go/src/filestore/types"
	tool_errors "github.com/elastifile/emanage-go/src/tools/errors"
	"github.com/elastifile/emanage-go/src/logging"
	"github.com/elastifile/emanage-go/src/types"
)

type ToolPropertiesBase struct {
	params  *types.ToolParams
	logger  log.Logger
	watcher chan error
}

type DefaultTool struct {
	Props types.ToolProperties
}

var Abort func(context *types.Context, errs ...error) error

func NewProperties(params *types.ToolParams) (result *ToolPropertiesBase, err error) {
	if len(params.TargetLoaders) < 1 {
		return nil, tool_errors.ErrNoLoaders
	}
	return &ToolPropertiesBase{
		params:  params,
		logger:  logging.NewLogger(string(params.ToolName)),
		watcher: make(chan error),
	}, nil
}

type ParseOpts struct {
	SourceBucket filestore_types.Bucket
	ConfigFile   string
	Template     string
	Data         interface{}
}

func (t *ToolPropertiesBase) ParseTemplate(opts *ParseOpts) (string, error) {
	var templateFile *types.NamedReader
	conf := agent.Config()

	client, err := filestore.New(conf.FilestoreHostInternal(), false)
	if err != nil {
		return "", err
	}

	if opts.ConfigFile != "" {
		files, e := client.ListBucket(&filestore.ListBucketOpts{
			Bucket:   opts.SourceBucket,
			Patterns: []string{opts.ConfigFile},
		})
		if e != nil {
			return "", e
		}
		if len(files) == 0 {
			return "", fmt.Errorf("specified template name '%v', but file was not provided nor found under bucket: %s",
				opts.ConfigFile, opts.SourceBucket)
		}

		// Download config file from minio
		objs, e := client.Download(opts.SourceBucket, opts.ConfigFile)
		if e != nil || len(objs) == 0 {
			return "", fmt.Errorf("specified template name '%v', but file was not provided, err: %v",
				opts.ConfigFile, e)
		}

		templateFile = objs[0]
		t.Logger().Debug("Using user-supplied template file", "configFile", templateFile.Name)

	} else {
		opts.ConfigFile = fmt.Sprintf("default_config_file_%v", t.Name())

		templateFile = types.NewNamedReader(opts.ConfigFile, bytes.NewReader([]byte(opts.Template)))
		t.Logger().Debug("Using default template", "configFile", opts.ConfigFile)
	}

	if opts.Data == nil {
		opts.Data = t.Params().Config
	}

	// parse by template
	parsed, err := textutil.ParseTemplate(&textutil.ParseOpts{
		Data:     opts.Data,
		Template: templateFile,
		Title:    string(t.params.ToolName) + "Template",
	})
	if err != nil {
		return "", err
	}

	// Upload back the parsed file to minio
	parsedFile := types.NewNamedReader(opts.ConfigFile, bytes.NewReader([]byte(parsed)))
	if err = client.Upload(opts.SourceBucket, parsedFile); err != nil {
		return "", err
	}

	return opts.ConfigFile, nil
}

func (t *ToolPropertiesBase) Name() types.ToolName {
	return t.params.ToolName
}

func (t *ToolPropertiesBase) WaitFor() (hosts []types.Host) {
	return t.params.Config.Loaders()
}

func (t *ToolPropertiesBase) Params() *types.ToolParams {
	return t.params
}

func (t *ToolPropertiesBase) Logger() log.Logger {
	return t.logger
}

func (t *ToolPropertiesBase) ErrorWatcher() chan error {
	return t.watcher
}

func (t *ToolPropertiesBase) NewLogAggregator(c *types.Container) types.LogAggregator {
	return NewNopAggregator()
}

func (t *ToolPropertiesBase) Logs() map[types.Host]*types.NamedReader {
	return nil
}

func ValidateConfigFiles(bucket filestore_types.Bucket, files []string) error {
	conf := agent.Config()

	client, err := filestore.New(conf.FilestoreHostInternal(), false)
	if err != nil {
		return err
	}

	logger.Debug("Validating config files in tool's bucket", "bucket", bucket, "files", files)
	for _, f := range files {
		if f == "" { //skip empty filenames (config file is omitted - not used)
			continue
		}
		objs, err := client.ListBucket(&filestore.ListBucketOpts{
			Bucket:   bucket,
			Patterns: []string{filepath.Base(f)},
		})
		if err != nil {
			return err
		}
		if len(objs) == 0 {
			return errors.Errorf("Validate config files: Failed to find config file '%s' at given bucket: %s.", f, bucket)
		}
	}
	return nil
}
