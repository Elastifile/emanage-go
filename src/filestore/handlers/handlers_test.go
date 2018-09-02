package filestore_handlers

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	minio_go "github.com/minio/minio-go"
	gt "common/gotest"

	"github.com/elastifile/emanage-go/src/filestore"
	"github.com/elastifile/emanage-go/src/logging"
	"github.com/elastifile/emanage-go/src/filestore/types"
	filestore_types "github.com/elastifile/emanage-go/src/filestore/types"
	logging_config "github.com/elastifile/emanage-go/src/logging/config"
)

func init() {
	logging.Setup(logging_config.ConfigForUnitTest())
}

func TestFilterCachedObjects(t *testing.T) {
	obj := minio_go.ObjectInfo{
		Key:          "testFile",
		LastModified: time.Now(),
	}

	// Found non-existing new file
	//////////////////////////////
	t.Logf("Expecting to find a new file ...")
	objects := filterCachedObjects([]minio_go.ObjectInfo{obj})
	if len(objects) == 0 {
		t.Fatalf("Expected to have at least one file (obj: %v), len(files) == 0", obj.Key)
	}
	if objects[0].Key != obj.Key {
		t.Fatalf("Got file %v while expected: %v)", objects[0].Key, obj.Key)
	}

	// Found exactly same file
	//////////////////////////
	t.Logf("Expecting to receive empty file list (there was no change)")
	objects = filterCachedObjects([]minio_go.ObjectInfo{obj})
	if len(objects) != 0 {
		t.Fatal("Expected not to be updated (obj.LastModified didn't change!), len(files) != 0")
	}

	// Found newer file
	///////////////////
	t.Logf("Expecting to find an updated file ...")
	obj.LastModified = time.Now().Add(time.Hour) // updating obj mTime

	objects = filterCachedObjects([]minio_go.ObjectInfo{obj})
	if len(objects) == 0 {
		t.Fatalf("Expected to have at least one file (obj.Key: %v), len(objects) == 0", obj.Key)
	}
	if objects[0].Key != obj.Key {
		t.Fatalf("Got file %v while expected: %v", objects[0].Key, obj.Key)
	}
}

//// Integration tests /////
func TestSyncMinio(t *testing.T) {
	gt.SkipIntegrationTest(t)

	// create bucket
	bucket := filestore_types.Bucket("minio-sync-test")
	path := bucket.WithDataVolume().String() // will usually be /data/<bucket name>/...
	// requires permission to create /data if not exists
	gt.Check(t, os.MkdirAll(path, 0755))

	filename := "foo"
	t.Log("Creating temp file ...")
	_, err := createTempFile(path, filename)
	gt.Check(t, err)

	minioEndpoint := types.Host(gt.Server(t))
	client, err := filestore.New(minioEndpoint, false)
	gt.Check(t, err)
	err = ensureBuckets(client, bucket)
	gt.Check(t, err)

	t.Log("Syncing minio ...")
	currentHost = types.Host("localhost")
	err = syncMinio(minioEndpoint, bucket, false)
	gt.Check(t, err)

	// cleanup after test if passes
	_ = os.RemoveAll(path)
}

func createTempFile(dir, fileName string) (*os.File, error) {
	text := fmt.Sprintf("*** Test file example - created on: %v ***\n",
		time.Now().Local())

	if err := ensurePath(dir); err != nil {
		return nil, err
	}

	f, err := ioutil.TempFile(dir, fileName)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(f.Name(), []byte(text), os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
