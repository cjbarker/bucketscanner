package bucketscanner

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Variables to identify the build
var (
	Version string
	Build   string
)

const bucketName string = "[replace-bucket-name]"

type BucketState int

// Bucket states
const (
	Unknown     BucketState = iota
	Invalid                 // Bucket does not exists e.g. 404 Not Found
	Private                 // Bucket exists but is not accessible e.g. 403 Forbidden
	Public                  // Bucket exists and is available e.g. 200 OK
	RateLimited             // Unable to determine due to rate limiting e.g. 503 Slow Down
)

type Scanner interface {
	Get(name string) (bucket *Bucket, err error)
	GetProviderName() (cloudProviderName string)
}

type Bucket struct {
	Provider  string      `json:"provider"`
	Name      string      `json:"name"`
	Scanned   time.Time   `json:"scanned"`
	URI       string      `json:"uri"`
	State     BucketState `json:"state"`
	NoFiles   int64       `json:"noFiles"`
	noDirs    int64
	TotalSize int64  `json:"totalSize"`
	Files     []file `json:"files"`
}

type file struct {
	Name  string `json:"name"`
	IsDir bool   `json:"directory"`
	Size  int64  `json:"size"`
	Files []file `json:"files"`
}

func getHTTPBucket(uri string) (contents *string, err error) {
	if strings.Trim(uri, " ") == "" {
		return nil, errors.New("Blank strings not accepted for bucket URI")
	}
	_, err = url.Parse(uri)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(uri)
	if resp.StatusCode != 200 {
		return nil, errors.New("Failed to get valid HTTP response due to STATUS code: " + resp.Status)
	}

	// only grab valid response
	defer resp.Body.Close()
	contentBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentStr := string(contentBytes)
	return &contentStr, nil
}
