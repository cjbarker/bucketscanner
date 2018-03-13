package bucketscanner

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strings"
	"time"
)

// Variables to identify the build
var (
	Version string
	Build   string
)

const bucketName string = "[replace-bucket-name]"

// BucketState denotes an integer defining its given state
type BucketState int

// Bucket states
const (
	Unknown     BucketState = iota
	Invalid                 // Bucket does not exists e.g. 404 Not Found
	Private                 // Bucket exists but is not accessible e.g. 403 Forbidden
	Public                  // Bucket exists and is available e.g. 200 OK
	RateLimited             // Unable to determine due to rate limiting e.g. 503 Slow Down
)

// Scanner interface declares functions for cloud provider scanner to implement
type Scanner interface {
	Read(name string) (bucket *Bucket, err error)
	Write(name string) (isWritable bool, err error)
	GetProviderName() (cloudProviderName string)
}

// Bucket structure is the results of a given bucket including its meta-data
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

// file is a representation of a bucket (object) file
type file struct {
	Name  string `json:"name"`
	IsDir bool   `json:"directory"`
	Size  int64  `json:"size"`
	Files []file `json:"files"`
	Body  []byte
}

// writeToArchive provides recursive HTTP bucket file download and writes to a given archive writer
func (b Bucket) writeToArchive(bucketFile *file, zipWriter *zip.Writer) (err error) {
	if &b == nil {
		return errors.New("Nil bucket unable to write to archive")
	}
	if bucketFile == nil || zipWriter == nil {
		return errors.New("Nil file and/or archive writer passed - unable to write to archive")
	}

	zipFile, err := zipWriter.Create(bucketFile.Name)
	if err != nil {
		return errors.New("Failed to create file in archive " + err.Error())
	}

	if len(bucketFile.Body) <= 0 {
		strBody, err := getHTTPBucket(b.URI + "/" + bucketFile.Name)
		if err != nil {
			return err
		}
		bucketFile.Body = []byte(*strBody)
	}

	_, err = zipFile.Write(bucketFile.Body)
	if err != nil {
		return errors.New("Failed to write file to archive " + err.Error())
	}

	return
}

// Download the contents of the bucket to a given destination directory
func (b Bucket) Download(destDir string) (archivePath *string, err error) {

	// set destination archive path to user's home dir
	if strings.Trim(destDir, " ") == "" {
		//return nil, errors.New("Destination directory is not accepted as a blank string")
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		destDir = usr.HomeDir + string(os.PathSeparator)
	}

	fi, err := os.Lstat(destDir)
	if os.IsNotExist(err) {
		return nil, errors.New("Destination file does NOT exist at " + destDir)
	}

	mode := fi.Mode()
	if !mode.IsDir() {
		return nil, errors.New("Destination file is NOT a directory " + destDir)
	}

	if len(b.Files) <= 0 {
		return nil, errors.New("Bucket " + b.Name + " has no files to download")
	}

	// Create buffer to write to archive
	t := time.Now()
	output := destDir + "bucket-" + b.Name + "-" + t.Format(time.RFC3339) + ".zip"
	newfile, err := os.Create(output)
	if err != nil {
		return nil, err
	}
	defer newfile.Close()

	// create zip archive
	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	// Iterate and download bucket files to the archive
	for _, file := range b.Files {
		err = b.writeToArchive(&file, zipWriter)
		if err != nil {
			return nil, errors.New("Failed to write file [" + file.Name + "] to archive: " + err.Error())
		}
	}

	return &output, nil
}

// getHTTPBucket establiesh HTTP connection to the uri and returns contents from HTTP response body
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
	contentBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	contentStr := string(contentBytes)
	return &contentStr, nil
}
