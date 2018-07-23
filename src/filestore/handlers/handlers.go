package filestore_handlers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go"

	"github.com/go-errors/errors"
	minio_go "github.com/minio/minio-go"

	filestore_types "filestore/types"

	"agent"
	"filestore"
	"logging"
	"messaging"
	"types"
	"runtime/debug"
)

var logger = logging.NewLogger("filestore_handlers")

const dirPerms = 0755

var _ = agent.RegisterAll(messaging.SubjectFilestoreSync, func(_, replySubject string, request messaging.FilestoreSyncRequest) {
	go func() {
		conf := agent.Config()
		logger.Info("Received request to sync slave with filestore", "host", conf.FilestoreHostInternal(), "request", request)
		err := syncMinio(conf.FilestoreHostInternal(), request.Bucket, false)
		agent.Publish(replySubject, messaging.RemoteDownloadResponse{
			BasicResponse: messaging.NewBasicResponse(err),
		})
	}()
})

// the following is at global scope to allow 'syncMinio' be tested (TestSyncMinio).
var currentHost types.Host

func syncMinio(filestoreEndpoint types.Host, srcBucket filestore_types.Bucket, verbose bool) error {
	// might be empty during integration test in which we'll set 'currentHost' manually through the test.
	if agent.Config().System.Setup.CurrentLoader != "" {
		currentHost = types.Host(agent.Config().System.CurrentLoaderInternal())
	}

	client, err := filestore.New(filestoreEndpoint, verbose)
	if err != nil {
		return err
	}

	// ensure local path exists
	localPath := srcBucket.ToDirName().WithDataVolume().String()
	logger.Debug("Sync filestore: local path", "localPath", localPath, "currentHost", currentHost)
	if err = ensurePath(localPath); err != nil {
		return err
	}

	// get upload list before downloading anything
	upList, err := lsDir(localPath)
	if err != nil {
		return err
	}

	// get download list
	objects, retries := []minio.ObjectInfo{}, 3
	logger.Debug("List bucket with retries until found objects [workaround for TESLA-2282]", "retries", retries)
	for i := 0; i < retries && len(objects) == 0; i++ {
		time.Sleep(500 * time.Millisecond) // to reduce a race chance
		objects, err = client.ListBucket(&filestore.ListBucketOpts{Bucket: srcBucket})
		if err != nil {
			return errors.Errorf("%v (bucket: %v, endpoint: %s)", err, srcBucket, filestoreEndpoint)
		}
	}

	// filter download list (download only new/updated objects)
	filtered := filterCachedObjects(objects)

	for _, o := range filtered {
		filename := o.Key
		fullFilename := filepath.Join(localPath, filename)

		logger.Debug("Will download file (syncMinio)", "file", filename, "target path", localPath)
		if err := client.Client.FGetObject(srcBucket.String(), filename, fullFilename); err != nil {
			return errors.Wrap(err, 0)
		}
		logger.Info("Downloaded file", "file", fullFilename)

		mode := os.ModePerm
		if err = os.Chmod(fullFilename, mode); err != nil {
			return errors.Wrap(err, 0)
		}
		logger.Debug("chmod file", "mode", mode, "file", fullFilename)
	}

	//// upload
	if len(upList) == 0 { // nothing to upload
		return nil
	}

	// ensure upload bucket exists
	if currentHost == "" {
		logger.Debug("Note: syncMinio currentHost is empty", "Stack", string(debug.Stack()))
	}
	dstBucket := srcBucket.ByHost(string(currentHost))
	logger.Debug("Sync filestore: destination bucket", "dstBucket", dstBucket)
	if err = ensureBuckets(client, srcBucket, dstBucket); err != nil {
		return err
	}

	uploaded, err := client.UploadFiles(&filestore.TransferFilesOpts{
		Patterns:  upList,
		Bucket:    dstBucket,
		Path:      localPath,
		Recursive: true,
	})
	if err != nil {
		return errors.Errorf("%v (bucket: %v)", err, dstBucket)
	}
	logger.Debug("filestore handler", "uploaded files", uploaded)
	return nil
}

func ensurePath(path string) error {
	// ensure local path exists (e.g. "/data/12345/erun")
	err := os.MkdirAll(path, dirPerms)
	if err != nil {
		return err
	}
	return nil
}

func ensureBuckets(client *filestore.Filestore, buckets ...filestore_types.Bucket) error {
	// ensure bucket exists (e.g. "12345.erun.loader1")
	for _, b := range buckets {
		err := client.MakeBucket(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func lsDir(path string) (files []string, err error) {
	f, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range f {
		files = append(files, file.Name())
	}
	return
}

var objectsCacheMap = map[string]time.Time{}

// 'shouldDownload' checks which object needs to be updated
// by checking/comparing objects last modified time in 'objectsCacheMap'.
func filterCachedObjects(objects []minio_go.ObjectInfo) (filtered []minio_go.ObjectInfo) {
	for _, o := range objects {
		currentMtime, ok := objectsCacheMap[o.Key]
		if !ok { // new object found, should download
			logger.Debug("Found new object to download", "obj", o, "etag", o.ETag)
			objectsCacheMap[o.Key] = o.LastModified
			filtered = append(filtered, o)
			continue
		}
		if currentMtime.Before(o.LastModified) {
			// updated object found, should download
			filtered = append(filtered, o)
		}
	}
	return
}
