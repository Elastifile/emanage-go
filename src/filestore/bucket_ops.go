package filestore

import (
	"path/filepath"

	"github.com/minio/minio-go"

	filestore_types "filestore/types"
)

func (m *Filestore) MakeBucket(bucket filestore_types.Bucket) error {
	bucketName := bucket.ToBucketName().String()
	if err := m.Client.BucketExists(bucketName); err == nil {
		// nothing to do, already exists
		return nil
	}
	location := ""                                    // we don't care about location
	logger.Warn("MMMMM MakeBucket", "bucket", bucket) //TODO: DELME
	return m.Client.MakeBucket(bucketName, location)
}

func (m *Filestore) RemoveBucket(bucket filestore_types.Bucket) error {
	objs, err := m.ListBucket(&ListBucketOpts{
		Bucket:   bucket,
		Patterns: []string{},
	})
	if err != nil {
		return err
	}

	for _, o := range objs {
		err := m.Client.RemoveObject(bucket.String(), o.Key)
		if err != nil {
			return err
		}
	}

	return m.Client.RemoveBucket(bucket.String())
}

type ListBucketOpts struct {
	Bucket   filestore_types.Bucket
	Patterns []string
}

// Recursively list all objects in given bucket
func (m *Filestore) ListBucket(opts *ListBucketOpts) ([]minio.ObjectInfo, error) {
	if err := m.Client.BucketExists(opts.Bucket.String()); err != nil {
		return nil, err
	}

	if len(opts.Patterns) == 0 {
		logger.Debug("Got empty pattern list! bringing all ...")
		opts.Patterns = []string{"*"}
	}

	doneCh := make(chan struct{})
	defer close(doneCh)
	recursive := true
	objects := []minio.ObjectInfo{}
	bucketName := opts.Bucket.ToBucketName().String()
	for objChan := range m.Client.ListObjects(bucketName, "", recursive, doneCh) {
		m.trace(objChan)
		if objChan.Err != nil {
			return nil, objChan.Err
		}
		objects = append(objects, objChan)
	}

	filtered, err := filterObjects(objects, opts.Patterns)
	if err != nil {
		return nil, err
	}
	return filtered, nil
}

func filterObjects(objects []minio.ObjectInfo, patterns []string) (filtered []minio.ObjectInfo, err error) {
	for _, p := range patterns {
		logger.Debug("Matching pattern", "pattern", p)
		for _, obj := range objects {
			logger.Debug("Checking match", "pattern", p, "object name", obj.Key)
			matched, err := filepath.Match(p, obj.Key)
			if err != nil {
				return nil, err
			}
			if matched {
				logger.Debug("Found match", "file", obj.Key)
				filtered = append(filtered, obj)
			}
		}
	}
	return filtered, nil
}
