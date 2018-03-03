package bucketscanner

import (
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
	"time"
)

// Cloud Provider Bucket Constant
const (
	awsName = "Amazon Simple Storage Service (S3)"
	awsURI  = "https://" + bucketName + ".s3.amazonaws.com"
)

type AwsScanner struct {
	result ListBucketResult
}

type ListBucketResult struct {
	XMLName      xml.Name `xml:"ListBucketResult"`
	Name         string
	Prefix       string
	Marker       string
	MaxKeys      int
	IsTruncated  bool
	ContentsList []Contents `xml:"Contents"`
}

type Contents struct {
	XMLName      xml.Name `xml:"Contents"`
	Key          string
	LastModified string
	Etag         string `xml:"ETag"`
	Size         int
	StorageClass string
}

func (a AwsScanner) Get(name string) (bucket *Bucket, err error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("Blank strings not accepted for bucket name")
	}

	url := strings.Replace(awsURI, bucketName, name, 1)

	bucket = &Bucket{
		Provider: awsName,
		Name:     name,
		URI:      url,
		State:    Unknown,
		Scanned:  time.Now(),
	}

	var sleepMs int

	// Parse State
	for bucket.State == Unknown {
		// Head check before deeper analysis
		resp, err := http.Head(url)
		if err != nil {
			return nil, err
		}

		switch resp.StatusCode {
		case 200:
			bucket.State = Public
		case 403:
			bucket.State = Private
		case 404:
			bucket.State = Invalid
		case 503:
			sleepMs += 500
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			if sleepMs >= 10000 {
				bucket.State = RateLimited
			}
		default:
			bucket.State = Unknown
		}
	}

	// Retrieve available HTTP payload
	if bucket.State == Public {
		contents, err := getHTTPBucket(url)
		if err != nil {
			return nil, err
		}

		//fmt.Printf("Resp. Bucket Contents: %s\n", *contents)

		err = xml.Unmarshal([]byte(*contents), &a.result)
		if err != nil {
			return nil, err
		}

		/*
			fmt.Printf("Result: %s\n", a.result.Name)
			fmt.Printf("Result: %d\n", a.result.MaxKeys)
			fmt.Printf("Result: %t\n", a.result.IsTruncated)
		*/

		for _, element := range a.result.ContentsList {
			bucket.NoFiles++
			bucket.TotalSize += int64(element.Size)
			/*
				fmt.Printf("\nResult: %s\n", element.Key)
				fmt.Printf("Result: %s\n", element.LastModified)
				fmt.Printf("Result: %s\n", element.Etag)
				fmt.Printf("Result: %d\n", element.Size)
				fmt.Printf("Result: %s\n", element.StorageClass)
			*/

		}
	}

	return bucket, nil
}

func (a AwsScanner) GetProviderName() (cloudProviderName string) {
	return awsName
}
