package filestore

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"logging"
	"types"

	gt "common/gotest"
	filestore_types "filestore/types"
	logging_config "logging/config"
)

func init() {
	logging.Setup(logging_config.ConfigForUnitTest())
}

// this test is used to test and prove that minio go client ListObjects
// is behaving unexpectedly and can sometimes return empty object list.
// You might need to run this test more then once to witness it.
func TestUnstableListBuckets(t *testing.T) {
	gt.SkipIntegrationTest(t)

	c, err := New(types.Host(gt.Server(t)), false)
	gt.Check(t, err)

	// create bucket
	b, err := getNewBucketUT(c, "test-list-buckets")
	gt.Check(t, err)

	// upload file
	fullFilename := createTempFile(t, "testFile")
	defer func() { _ = os.Remove(fullFilename) }()
	gt.Check(t, c.uploadFile(b.String(), fullFilename))

	// list bucket objects multiple times
	routines := 20
	done := make(chan struct{}, routines)
	defer close(done)

	count := 0

	listObjectsRoutine := func() {
		objects := []string{}
		t.Log("listing objects..")

		doneCh := make(chan struct{})
		defer close(doneCh)
		for objChan := range c.Client.ListObjects(b.ToBucketName().String(), "", true, doneCh) {
			objects = append(objects, objChan.Key)
		}

		if len(objects) != 1 {
			t.Logf("len(objects) != 1, objects: %v\n", objects)
		} else {
			count++
			t.Logf("Objects: %v\n", objects)
		}
		done <- struct{}{}
	}

	for i := 0; i < routines; i++ {
		go listObjectsRoutine()
	}

	for i := 0; i < routines; i++ {
		<-done
	}
	t.Log("Done listing objects")
	if count != routines {
		t.Fatalf("ERROR: count objects != num of attempts(routines). count: %v, attempts: %v", count, routines)
	}
	t.Log("Found expected objects count")
}

// Simple test of upload/download
// 1. upload files
// 2. download it
func TestUploadDownload(t *testing.T) {
	gt.SkipIntegrationTest(t)

	server := gt.Server(t)

	c, err := New(types.Host(server), false)
	gt.Check(t, err)

	b := filestore_types.Bucket("test1")
	gt.Check(t, c.MakeBucket(b))

	filename := "uploadedTestFile"
	fullfilename := createTempFile(t, filename)
	defer func() { _ = os.Remove(fullfilename) }()

	t.Logf("Will upload file: %v\n", fullfilename)
	patterns := []string{fmt.Sprintf("%v*", filename)}

	files, err := c.UploadFiles(&TransferFilesOpts{
		Bucket:    b,
		Path:      fullfilename,
		Patterns:  patterns,
		Recursive: true,
	})
	gt.Check(t, err)
	if len(files) == 0 {
		t.Fatal("len(uploaded files) == 0")
	}

	fileHandles, err := c.Download(b, patterns...)
	gt.Check(t, err)
	if len(fileHandles) == 0 {
		t.Fatal("len(downloaded files) == 0")
	}
}

func getNewBucketUT(client *Filestore, name string) (filestore_types.Bucket, error) {
	b := filestore_types.Bucket(strings.ToLower(name))
	if err := client.Client.BucketExists(name); err == nil {
		err := client.RemoveBucket(b)
		if err != nil {
			return "", err
		}
	}
	err := client.MakeBucket(b)
	if err != nil {
		return "", err
	}
	return b, nil
}

func createTempFile(t *testing.T, fileName string) (filename string) {
	text := fmt.Sprintf("This test file created on: %v\n",
		time.Now().Local())

	tmpDir := "/tmp"
	f, err := ioutil.TempFile(tmpDir, fileName)
	gt.Check(t, err)

	err = ioutil.WriteFile(f.Name(), []byte(text), os.ModePerm)
	gt.Check(t, err)

	return f.Name()
}
