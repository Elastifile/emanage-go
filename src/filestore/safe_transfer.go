package filestore

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-errors/errors"

	filestore_types "github.com/elastifile/emanage-go/src/filestore/types"
	"github.com/elastifile/emanage-go/src/filestore/types"
)

type SafeTransferReport struct {
	Text  string
	Err   error
	Grace time.Duration
}

func RemoveBucket(host types.Host, bucketName string, timeout time.Duration, progress func(report *SafeTransferReport)) error {
	var (
		err           error
		fs            *Filestore
		fstoreVerbose bool = false
	)
	bucket := filestore_types.Bucket(bucketName)
	for start := time.Now(); time.Since(start) < timeout; time.Sleep(2 * time.Second) {
		fs, err = newFilestore(host, timeout-time.Since(start), fstoreVerbose, progress)
		if err != nil {
			progress(&SafeTransferReport{
				Text:  "Failed to create filestore client",
				Err:   err,
				Grace: roundDuration(timeout - time.Since(start)),
			})
		} else {
			err = fs.RemoveBucket(bucket)
			if err != nil {
				progress(&SafeTransferReport{
					Text:  "Failed to remove bucket: " + bucket.ToBucketName().String(),
					Err:   err,
					Grace: roundDuration(timeout - time.Since(start)),
				})
			} else {
				break
			}
		}
	}
	if err != nil && strings.Contains(err.Error(), "The specified bucket does not exist") {
		return nil
	}
	return err
}
func SafeUploadFiles(host types.Host, bucketName string, paths []string, timeout time.Duration, progress func(report *SafeTransferReport)) error {
	var (
		err           error
		fs            *Filestore
		fstoreVerbose bool = false
	)

	bucket := filestore_types.Bucket(bucketName)

	for start := time.Now(); time.Since(start) < timeout; time.Sleep(2 * time.Second) {
		fs, err = newFilestore(host, timeout-time.Since(start), fstoreVerbose, progress)
		if err != nil {
			progress(&SafeTransferReport{
				Text:  "Failed to create filestore client",
				Err:   err,
				Grace: roundDuration(timeout - time.Since(start)),
			})
		} else {
			if err = fs.MakeBucket(bucket); err != nil {
				progress(&SafeTransferReport{
					Text:  "Failed to make bucket: " + bucket.ToBucketName().String(),
					Err:   err,
					Grace: roundDuration(timeout - time.Since(start)),
				})
				continue
			}

			uploadMap := make(map[string][]string)
			for _, path := range paths {
				fileDir := filepath.Dir(path)
				fileName := filepath.Base(path)
				if _, exists := uploadMap[fileDir]; exists {
					uploadMap[fileDir] = append(uploadMap[fileDir], fileName)
				} else {
					uploadMap[fileDir] = []string{fileName}
				}
			}

			for dir, names := range uploadMap {
				_, err = fs.UploadFiles(&TransferFilesOpts{
					Bucket:    bucket,
					Path:      dir,
					Patterns:  names,
					Recursive: false,
				})
				if err != nil {
					progress(&SafeTransferReport{
						Text:  fmt.Sprintf("Failed uploading %s/%v to filestore", dir, names),
						Err:   err,
						Grace: roundDuration(timeout - time.Since(start)),
					})
					break
				}
			}
		}
		if err == nil {
			break
		}
	}

	return err
}

func SafeReadFiles(host types.Host, bucketName string, paths []string, timeout time.Duration, progress func(report *SafeTransferReport)) ([][]byte, error) {
	var (
		err           error
		fs            *Filestore
		fileNames     []string
		readers       []*types.NamedReader
		fstoreVerbose bool = false
		bodies        [][]byte
	)

	bucket := filestore_types.Bucket(bucketName).ToBucketName()

	for _, path := range paths {
		fileNames = append(fileNames, filepath.Base(path))
	}

	for start := time.Now(); time.Since(start) < timeout; time.Sleep(2 * time.Second) {

		fs, err = newFilestore(host, timeout-time.Since(start), fstoreVerbose, progress)
		if err != nil {
			progress(&SafeTransferReport{
				Text:  "Failed to create filestore client",
				Err:   err,
				Grace: roundDuration(timeout - time.Since(start)),
			})
		} else {
			readers, err = fs.Download(bucket, fileNames...)
			if err == nil && len(readers) > 0 {
				break
			} else {
				progress(&SafeTransferReport{
					Text:  "Failed downloading from filestore",
					Err:   err,
					Grace: roundDuration(timeout - time.Since(start)),
				})
			}
		}
	}

	if len(readers) == 0 {
		return nil, errors.Errorf("No downloads for files: %s, bucket: %s, host: %s", fileNames, bucket, host)
	}

	sort.Sort(types.ByNamedReaders(readers))
	for _, reader := range readers {
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		progress(&SafeTransferReport{
			Text: "Read body of file:" + string(reader.Name),
		})
		bodies = append(bodies, body)
	}

	progress(&SafeTransferReport{
		Text: fmt.Sprintf("Read body count: %d", len(bodies)),
	})
	return bodies, nil
}

func newFilestore(host types.Host, grace time.Duration, verbose bool, progress func(report *SafeTransferReport)) (*Filestore, error) {
	fs, err := New(host, verbose)
	if err != nil {
		progress(&SafeTransferReport{
			Text:  "Failed creating filestore client",
			Err:   err,
			Grace: grace,
		})
	} else {
		progress(&SafeTransferReport{
			Text: "Created filestore client created, host: " + string(host),
		})
	}
	return fs, err
}

func roundDuration(dur time.Duration) time.Duration {
	return dur - time.Duration(dur.Nanoseconds())
}
