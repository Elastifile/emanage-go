package types

import (
	"fmt"
	"testing"
)

func TestBucket(t *testing.T) {
	orig := "/data/db45463a5762/erun"
	fmt.Printf("original string: %v\n", orig)
	b := Bucket(orig)

	if b.String() != orig {
		t.Fatalf("Bucket name should be: %v (got %v)", orig, b)
	}

	woDataVolume := "db45463a5762/erun"
	b = b.TrimDataVolume()
	if b.String() != woDataVolume {
		t.Fatalf("Bucket name without DataVolume should be: %v (got %v)", woDataVolume, b)
	}

	bucketNameWoDataVolume := "db45463a5762.erun"
	b = b.TrimDataVolume().ToBucketName()
	if b.String() != bucketNameWoDataVolume {
		t.Fatalf("Bucket name without DataVolume should be: %v (got %v)", bucketNameWoDataVolume, b)
	}
}
