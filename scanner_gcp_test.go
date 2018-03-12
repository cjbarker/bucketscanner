package bucketscanner_test

import (
	"gitlab.com/cjbarker/bucketscanner"
	"testing"
)

func TestGetProviderName(t *testing.T) {
	expected := "Google Cloud Storage"
	gcp := &bucketscanner.GcpScanner{}
	if gcp.GetProviderName() != expected {
		t.Errorf("Invalid GCP provider name. got: %s, expected %s", gcp.GetProviderName(), expected)
	}
}

func TestWriteGcp(t *testing.T) {
	aws := &bucketscanner.GcpScanner{}
	_, err := aws.Write("   ")
	if err == nil {
		t.Errorf("Error should occur when empty bucket name string is attempted to be retrieved.")
	}
}
