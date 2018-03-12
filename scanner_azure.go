package bucketscanner

import (
	"errors"
	"strings"
)

// Cloud Provider Bucket Constant
// https://docs.microsoft.com/en-us/azure/storage/blobs/storage-dotnet-how-to-use-blobs
const (
	azureName = "Azure Blob Storage"
	azureURI  = "https://storagesample.blob.core.windows.net/" + bucketName
)

// AzureScanner is struct for cloud scanner of Azure
type AzureScanner struct {
}

// GetProviderName returns the given Cloud Provider's name for the scanner
func (a AzureScanner) GetProviderName() (cloudProviderName string) {
	return azureName
}

func (a AzureScanner) Read(name string) (bucket *Bucket, err error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("Blank strings not accepted for bucket name")
	}

	//	url := strings.Replace(azureURI, bucketName, name, 1)

	// TODO implement
	/*
		bucket = &Bucket{
			provider: g.Getprovider(),
			name:     name,
		}
		return bucket, nil
	*/

	return nil, errors.New("Azure Scanner is currently not supported")
}

func (a AzureScanner) Write(name string) (isWritable bool, err error) {
	if strings.Trim(name, " ") == "" {
		return false, errors.New("Blank strings not accepted for bucket name")
	}

	//url := strings.Replace(azureURI, bucketName, name, 1)

	// TODO implement

	return false, errors.New("Azure Writer is currently not supported")
}
