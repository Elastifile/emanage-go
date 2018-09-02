package filestore

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-errors/errors"
	minio "github.com/minio/minio-go"
	"github.com/hashicorp/go-multierror"

	filestore_types "github.com/elastifile/emanage-go/src/filestore/types"
	"github.com/elastifile/emanage-go/src/filestore/types"
)

type TransferFilesOpts struct {
	Bucket    filestore_types.Bucket
	Patterns  []string
	Path      string
	Recursive bool
}

func (m *Filestore) UploadFiles(opts *TransferFilesOpts) (loadedFiles []string, errRet error) {
	logger.Debug("Uploading files", "host", m.Endpoint, "bucket", opts.Bucket, "path", opts.Path, "patterns", opts.Patterns)

	if opts.Recursive {
		errRet = filepath.Walk(opts.Path, func(pathname string, info os.FileInfo, err error) error {
			if err != nil {
				logger.Error(err.Error())
				return nil
			}

			if info.IsDir() {
				return nil
			}

			if info.Size() == 0 {
				logger.Debug("Skipping upload of empty file (size == 0)", "file", pathname)
				return nil
			}

			if relname, err := m.match(opts.Patterns, opts.Path, pathname); err != nil {
				return err
			} else if relname == "" {
				logger.Debug("Skipping non-matching file", "path", pathname)
			} else {
				loadedFiles = append(loadedFiles, relname)
			}

			return nil
		})
	} else {
		for _, p := range opts.Patterns {
			if _, err := os.Stat(filepath.Join(opts.Path, p)); !os.IsNotExist(err) {
				loadedFiles = append(loadedFiles, p)
			}
		}
	}

	for _, relname := range loadedFiles {
		pathname := filepath.Join(opts.Path, relname)
		if err := m.uploadFile(opts.Bucket.String(), pathname); err != nil {
			return nil, err
		}
	}

	logger.Info("Uploaded files",
		"host", m.Endpoint,
		"bucket", opts.Bucket,
		"path", opts.Path,
		"patterns", opts.Patterns,
	)

	return loadedFiles, errRet
}

func (m *Filestore) match(patterns []string, path string, pathname string) (string, error) {
	if len(patterns) == 0 {
		logger.Debug("No patterns provided, matching any file")
		patterns = append(patterns, "*")
	}

	relname, err := filepath.Rel(path, pathname)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var matched bool
	for _, p := range patterns {
		logger.Debug("Checking match", "pattern", p, "name", relname)
		matched, err = filepath.Match(p, relname)
		if err != nil {
			return "", err
		}
		if matched {
			logger.Debug("Found match", "file", relname)
			break
		}
	}

	return relname, nil
}

func (m *Filestore) DownloadFiles(opts *TransferFilesOpts) (files []string, err error) {
	logger.Debug("Downloading files",
		"bucket", opts.Bucket, "path", opts.Path, "patterns", opts.Patterns)

	if len(opts.Patterns) == 0 {
		logger.Debug("No patterns provided, matching any file")
		opts.Patterns = append(opts.Patterns, "*")
	}

	if opts.Bucket == "" {
		return nil, errors.Errorf("No source bucket provided (opts: %v)", opts)
	}

	// get all bucket's objects
	objects, err := m.ListBucket(&ListBucketOpts{
		Bucket:   opts.Bucket,
		Patterns: opts.Patterns,
	})
	if err != nil {
		return nil, err
	}

	// get handle of matching objects
	for _, obj := range objects {
		filename := obj.Key
		err := m.downloadFile(opts.Bucket.String(), filename, opts.Path)
		if err != nil {
			return nil, err
		}
		err = os.Chmod(filepath.Join(opts.Path, filename), 0755)
		if err != nil {
			return nil, err
		}
		files = append(files, filename)

	}
	logger.Info("Downloaded files",
		"host", m.Endpoint,
		"bucket", opts.Bucket,
		"path", opts.Path,
		"files", files,
	)

	return files, nil
}

func (m *Filestore) Download(b filestore_types.Bucket, patterns ...string) (files []*types.NamedReader, err error) {
	logger.Info("Download", "bucket", b, "patterns", patterns, "endpoint", m.Endpoint)

	if len(patterns) == 0 {
		logger.Info("No patterns provided, matching any file")
		patterns = append(patterns, "*")
	}

	if b == "" {
		return nil, errors.Errorf("No source bucket provided")
	}

	// get all bucket's objects
	listOpts := ListBucketOpts{
		Bucket:   b,
		Patterns: patterns,
	}
	objects, err := m.ListBucket(&listOpts)
	if err != nil {
		return nil, err
	}
	logger.Info("list bucket", "opts", listOpts, "objects", len(objects))

	// get reader of matching objects
	for _, obj := range objects {
		r, e := m.Client.GetObject(b.String(), obj.Key)
		if e != nil {
			e = errors.WrapPrefix(e, fmt.Sprintf("Failed to read object %v from bucket %v", obj.Key, b.String()), 0)
			err = multierror.Append(err, e)
			logger.Debug("Failed to get object", "err", e)
			continue
		}
		files = append(files, types.NewNamedReader(obj.Key, r))
		logger.Debug("loaded bucket object", "key", obj.Key)
	}
	return files, err
}

func (m *Filestore) Upload(b filestore_types.Bucket, r *types.NamedReader) (err error) {
	logger.Debug("Minio: Uploading ...", "bucket", b, "fileName", r.Name)
	contentType := ""
	n, err := m.Client.PutObject(
		b.ToBucketName().String(),
		filepath.Base(r.Name),
		r,
		contentType,
	)
	if err != nil {
		return err
	}
	logger.Debug(fmt.Sprintf("upload: wrote %v bytes", n))
	return nil
}

func (m *Filestore) uploadFile(bucket string, file string) (err error) {
	logger.Debug("Uploading file", "bucket", bucket, "file", file)
	contentType := ""
	n, err := m.Client.FPutObject(
		bucket,
		filepath.Base(file),
		file,
		contentType)
	if err != nil {
		logger.Error("uploadFile failed",
			"bucket", bucket,
			"file name", filepath.Base(file),
			"file path", file,
			"contentType", contentType)
		return err
	}
	logger.Debug("uploaded", "file", file, "bytes", n)
	return nil
}

func (m *Filestore) downloadFile(bucket string, fileName string, downloadPath string) (err error) {
	logger.Debug("Downloading file from filestore",
		"bucket", bucket,
		"fileName", fileName,
		"target path", downloadPath)

	return m.Client.FGetObject(
		bucket,
		fileName,
		filepath.Join(downloadPath, fileName),
	)
}

func (m *Filestore) CopyBucket(srcBucket, dstBucket filestore_types.Bucket) error {
	objs, err := m.ListBucket(&ListBucketOpts{
		Bucket:   srcBucket,
		Patterns: []string{},
	})
	if err != nil {
		return err
	}

	if err := m.MakeBucket(dstBucket); err != nil {
		return err
	}

	logger.Debug("Copying source bucket content to destination bucket", "srcBucket", srcBucket, "dstBucket", dstBucket)
	for _, obj := range objs {
		src := fmt.Sprintf("%v/%v", srcBucket, obj.Key)
		err := m.CopyObject(dstBucket.String(), obj.Key, src)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Filestore) CopyObject(bucketName string, objectName string, objectSource string) (err error) {
	return m.Client.CopyObject(bucketName, objectName, objectSource, minio.CopyConditions{})
}
