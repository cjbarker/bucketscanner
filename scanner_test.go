package bucketscanner_test

import (
	"gitlab.com/cjbarker/bucketscanner"
	"os"
	"testing"
)

func TestBuild(t *testing.T) {
	if len(bucketscanner.Version) <= 0 {
		t.Errorf("Version string length was empty, zero or less; got: %s", bucketscanner.Version)
	}
	if len(bucketscanner.Build) <= 0 {
		t.Errorf("Build string length was empty, zero or less; got: %s", bucketscanner.Build)
	}
}

func TestDownload(t *testing.T) {
	aws := &bucketscanner.AwsScanner{}

	bucket, err := aws.Read(PublicBucket)
	if err != nil {
		t.Errorf("Was expecting bucket %s to provide Public state, but got error: %s", PublicBucket, err.Error())
	}

	_, err = bucket.Download("    ")
	if err == nil {
		t.Errorf("Should not be able to download bucket to empty directory filename")
	}

	// create test regular file
	tmpFilename := ".test_temp_file"
	dst, err := os.Create(tmpFilename)
	if err != nil {
		t.Errorf("Unable to create test file due to error: %s", err.Error())
	}

	defer dst.Close()

	_, err = bucket.Download(tmpFilename)
	if err == nil {
		t.Errorf("Should not be able to download bucket to regular file: %s", tmpFilename)
	}

	os.Remove(tmpFilename)

	// create test directory
	tmpFilename = os.TempDir()
	_, err = bucket.Download(tmpFilename)
	if err != nil {
		t.Errorf("Unable to download bucket to dir due to error %s", err.Error())
	}
}
