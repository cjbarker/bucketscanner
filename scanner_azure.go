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

type AzureScanner struct {
}

func (a AzureScanner) GetProviderName() (cloudProviderName string) {
	return azureName
}

func (a AzureScanner) Get(name string) (bucket *Bucket, err error) {
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
