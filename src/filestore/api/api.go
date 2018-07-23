package filestore_api

import (
	"fmt"
	"time"

	"github.com/go-errors/errors"

	"filestore"
	filestore_types "filestore/types"
	"logging"
	"messaging"
	"types"
)

var logger = logging.NewLogger("filestore_api")

const dirPerms = 0755

type client struct {
	msg       messaging.Client
	filestore *filestore.Filestore
}

// filestoreEndpoint = minio server's dns reachable name or ip
// verbose       = minio http requests verbosity
func NewClient(msg messaging.Client, filestoreEndpoint types.Host, verbose bool) (*client, error) {
	m, err := filestore.New(filestoreEndpoint, verbose)
	if err != nil {
		return nil, err
	}

	return &client{
		msg:       msg,
		filestore: m,
	}, nil
}

// CCStore the provided files to the remote hosts.
func (c *client) RequestStoreFiles(bucket filestore_types.Bucket, path string, hosts []types.Host) error {
	logger.Debug("Request store files", "path", path, "bucket", bucket, "hosts", hosts)
	// tell all hosts to sync with given bucket
	return c.sendSync(bucket, hosts)
}

// Sync matching patterns files from remote hosts to minio and download them here.
func (c *client) RequestRetrieveFiles(path string, filePatterns []string, hosts []types.Host) (filesByHost types.FilesByHost, err error) {
	bucket := filestore_types.Bucket(path).TrimDataVolume().ToBucketName()

	logger.Debug("Request retrieve files", "path", path, "bucket", bucket, "patterns", filePatterns, "hosts", hosts)

	// 1. tell all hosts to sync with filestore
	err = c.sendSync(bucket, hosts)
	if err != nil {
		return nil, err
	}

	filesByHost = make(types.FilesByHost)
	for _, h := range hosts {
		// 2. download from filestore
		b := bucket.ByHost(string(h))
		logger.Debug("download files per host", "bucket", b, "patterns", filePatterns)

		if err := c.filestore.Client.BucketExists(b.String()); err != nil {
			// bucket doesn't exists, we'll set the host key value in map with empty file slice.
			filesByHost[h] = make([]*types.NamedReader, 0)
			continue
		}

		files, err := c.filestore.Download(b, filePatterns...)
		if err != nil {
			return nil, errors.Errorf("RequestRetrieveFiles: Failed download from bucket: %v (err: %v,  filePatterns: %v)", b, err, filePatterns)
		}

		logger.Debug("got files", "files", toFileNames(files))
		filesByHost[h] = files
	}

	logger.Debug("The following files were downloaded", "files", filesByHost)
	return filesByHost, nil
}

// sendSync will broadcast a filestore sync request to all registered handlers on subject 'SubjectFilestoreSync'.
// all registered handlers will download and upload files to the given bucket.
// download/upload path is determined from the bucket name, Example:
// bucket '77c7251b1577/erun' will download/upload from '/data/77c7251b1577/erun')
func (c *client) sendSync(bucket filestore_types.Bucket, hosts []types.Host) error {
	request := messaging.FilestoreSyncRequest{
		BasicRequest: messaging.NewBasicRequest(),
		Bucket:       bucket,
	}
	request.Context.Timeout = 5 * time.Minute
	logger.Debug("Sending filestore sync request", "request", fmt.Sprintf("%+v", request))

	fileResponses := make([]messaging.RemoteDownloadResponse, len(hosts))
	responseRefs := make([]messaging.Response, len(fileResponses))

	for i := range responseRefs {
		responseRefs[i] = &fileResponses[i]
	}

	return c.msg.RequestWithMultiResponse(messaging.SubjectFilestoreSync, &request, responseRefs)
}

func toFileNames(files []*types.NamedReader) (names []string) {
	for _, f := range files {
		names = append(names, f.Name)
	}
	return
}
