package bucketscanner

import (
	"errors"
	"strings"
)

// Cloud Provider Bucket Constant
// https://cloud.google.com/storage/docs/xml-api/reference-uris
const (
	gcpName = "Google Cloud Storage"
	gcpURI  = "https://" + bucketName + ".storage.googleapis.com"
)

type GcpScanner struct {
}

func (g GcpScanner) GetProviderName() (cloudProviderName string) {
	return gcpName
}

func (g GcpScanner) Read(name string) (bucket *Bucket, err error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("Blank strings not accepted for bucket name")
	}

	//	url := strings.Replace(gcpURI, bucketName, name, 1)

	// TODO implement
	/*
		bucket = &Bucket{
			provider: g.Getprovider(),
			name:     name,
		}
		return bucket, nil
	*/

	return nil, errors.New("GcpScanner is currently not supported")
}

func (g GcpScanner) Write(name string) (isWritable bool, err error) {
	if strings.Trim(name, " ") == "" {
		return false, errors.New("Blank strings not accepted for bucket name")
	}

	//url := strings.Replace(azureURI, bucketName, name, 1)

	// TODO implement

	return false, errors.New("GcpWWriter is currently not supported")
}
