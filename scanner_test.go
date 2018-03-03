package bucketscanner_test

import (
	"encoding/json"
	"gitlab.com/cjbarker/bucketscanner"
	"testing"
)

const (
	InvalidBucket = "espnasdfadsfasd"
	PrivateBucket = "espn"
	PublicBucket  = "listing-test"
)

func TestBuild(t *testing.T) {
	if len(bucketscanner.Version) <= 0 {
		t.Errorf("Version string length was empty, zero or less; got: %s", bucketscanner.Version)
	}
	if len(bucketscanner.Build) <= 0 {
		t.Errorf("Build string length was empty, zero or less; got: %s", bucketscanner.Build)
	}
}

func TestGetProviderName(t *testing.T) {
	var expected = "Amazon Simple Storage Service (S3)"
	aws := &bucketscanner.AwsScanner{}
	if aws.GetProviderName() != expected {
		t.Errorf("Invalid AWS provider name. got: %s, expected %s", aws.GetProviderName(), expected)
	}
}

func TestGetAws(t *testing.T) {
	// Empty bucket name
	aws := &bucketscanner.AwsScanner{}
	_, err := aws.Get("   ")
	if err == nil {
		t.Errorf("Error should occur when empty bucket name string is attempted to be retrieved.")
	}

	// Invalid bucket
	bucket, err := aws.Get(InvalidBucket)
	if err != nil {
		t.Errorf("Was expecting bucket %s to provide Invalid state, but got error: %s", InvalidBucket, err.Error())
	}
	if bucket.State != bucketscanner.Invalid {
		t.Errorf("Bucket state error, got: %d, expected %d", bucket.State, bucketscanner.Invalid)
	}

	// Private bucket
	bucket, err = aws.Get(PrivateBucket)
	if err != nil {
		t.Errorf("Was expecting bucket %s to provide Private state, but got error: %s", PrivateBucket, err.Error())
	}
	if bucket.State != bucketscanner.Private {
		t.Errorf("Bucket state error, got: %d, expected %d", bucket.State, bucketscanner.Private)
	}

	// Public bucket
	bucket, err = aws.Get(PublicBucket)
	if err != nil {
		t.Errorf("Was expecting bucket %s to provide Public state, but got error: %s", PublicBucket, err.Error())
	}
	if bucket.State != bucketscanner.Public {
		t.Errorf("Bucket state error, got: %d, expected %d", bucket.State, bucketscanner.Public)
	}

	// Valid JSON
	jsonStr, err := json.Marshal(bucket)
	if err != nil {
		t.Errorf("Unable to marshall bucket to JSON due to error: %s", err.Error())
	}
	if len(string(jsonStr)) < 10 {
		t.Errorf("Bucket JSON string is too short in length: %s", string(jsonStr))
	}
}
