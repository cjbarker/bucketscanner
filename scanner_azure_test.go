package bucketscanner_test

import (
	"gitlab.com/cjbarker/bucketscanner"
	"testing"
)

func TestGetAzureProviderName(t *testing.T) {
	expected := "Azure Blob Storage"
	azure := &bucketscanner.AzureScanner{}
	if azure.GetProviderName() != expected {
		t.Errorf("Invalid Azure provider name. got: %s, expected %s", azure.GetProviderName(), expected)
	}
}

func TestWriteAzure(t *testing.T) {
	azure := &bucketscanner.AzureScanner{}
	_, err := azure.Write("   ")
	if err == nil {
		t.Errorf("Error should occur when empty bucket name string is attempted to be retrieved.")
	}
}
