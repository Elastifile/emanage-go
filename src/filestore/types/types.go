package types

import (
	"fmt"
	"path/filepath"
	"strings"
)

// A valid Bucket name: (from http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html)
// - must be at least 3 and no more than 63 characters long.
// - must be a series of one or more labels, can be separated by a single period (.).
// - can contain lowercase letters, numbers, and hyphens.
// - each label must start and end with a lowercase letter or a number.
// - cannot start nor end with a period (just as a seperator).
type Bucket string

func (b Bucket) String() string { return string(b) }

const (
	separator  = "."
	slash      = "/"
	dataVolume = "/data/"
)

func (b Bucket) ToBucketName() Bucket {
	// similar to filepath.FromSlash, except we're using minio bucket seperator
	return Bucket(strings.Replace(b.String(), slash, separator, -1))
}

func (b Bucket) ToDirName() Bucket {
	// similar to filepath.ToSlash, except we're using minio bucket seperator
	return Bucket(strings.Replace(b.String(), separator, slash, -1))
}

func (b Bucket) TrimDataVolume() Bucket {
	//  Example: /data/77c7251b1577/erun ==> /77c7251b1577/erun
	return Bucket(strings.TrimPrefix(b.String(), dataVolume))
}

func (b Bucket) WithDataVolume() Bucket {
	//  Example: 77c7251b1577/erun ==> /data/77c7251b1577/erun
	return Bucket(filepath.Join("/data", b.ToDirName().String()))
}

func (b Bucket) ByHost(host string) Bucket {
	//  Example: 77c7251b1577/erun ==> 77c7251b1577.erun.loader1
	if host == "" {
		fmt.Println("WARN: Received empty host")
	}
	parts := []string{b.ToBucketName().String(), host}
	return Bucket(toLegalBucketName(strings.Join(parts, ".")))
}

func toLegalBucketName(name string) string {
	name = strings.ToLower(name)
	name = strings.Replace(name, "_", "-", -1)
	return name
}
