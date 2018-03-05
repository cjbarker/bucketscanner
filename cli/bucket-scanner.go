package main

import (
	"encoding/json"
	"fmt"
	"gitlab.com/cjbarker/bucketscanner"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
	"sync"
	"time"
)

// Exit Codes
const (
	Success      = 0
	InvalidCloud = 100
)

// Provider
const (
	All           = "all"
	AwsProvider   = "aws"
	GcpProvider   = "gcp"
	AzureProvider = "azure"
)

type Config struct {
	BucketNames   *string
	Download      *bool
	Verbose       *bool
	CloudProvider *string
	ThrottleMs    *int
	Json          *bool
}

func (c Config) v(msg string) {
	if *c.Verbose {
		fmt.Printf("%s\n", msg)
	}
}

// Globals
var configPtr *Config

func getScanner(providerName *string) (scanners []bucketscanner.Scanner) {
	if providerName == nil || strings.Trim(*providerName, " ") == "" {
		return nil
	}

	//var scanners []*Scanner
	if strings.ToLower(*providerName) == All {
		scanners = append(scanners, &bucketscanner.AwsScanner{})
		scanners = append(scanners, &bucketscanner.GcpScanner{})
		scanners = append(scanners, &bucketscanner.AzureScanner{})
	} else if strings.ToLower(*providerName) == AwsProvider {
		scanners = append(scanners, &bucketscanner.AwsScanner{})
	} else if strings.ToLower(*providerName) == GcpProvider {
		scanners = append(scanners, &bucketscanner.GcpScanner{})
	} else if strings.ToLower(*providerName) == AzureProvider {
		scanners = append(scanners, &bucketscanner.AzureScanner{})
	} else {
		scanners = nil
	}

	return scanners
}

func main() {
	configPtr = new(Config)

	app := kingpin.New(os.Args[0], "Cloud command-line bucket (object) scanner.")
	app.Version("Version: " + bucketscanner.Version + "\nBuild: " + bucketscanner.Build)

	configPtr.BucketNames = app.Arg("bucket-name", "Bucket(s) name(s) to scan. Does support comma separated for multiple buckets.").Required().String()
	configPtr.CloudProvider = app.Flag("cloud", "Cloud provider to scan: aws, gcp, azure. Defaults to all.").Required().String()
	configPtr.ThrottleMs = app.Flag("throttle", "Time in milliseconds to throttle subsequent requests sent to a given provider.").Int()
	configPtr.Download = app.Flag("download", "Download bucket content(s).").Bool()
	configPtr.Json = app.Flag("json", "Output results in JSON.").Bool()
	configPtr.Verbose = app.Flag("verbose", "Verbose output messages. Defaults to quiet.").Bool()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Default to all
	if strings.Trim(*configPtr.CloudProvider, " ") == "" {
		*configPtr.CloudProvider = All
	}

	// output settings
	configPtr.v(fmt.Sprintf("Cloud: %s", *configPtr.CloudProvider))
	configPtr.v(fmt.Sprintf("Buckets: % s", *configPtr.BucketNames))
	configPtr.v(fmt.Sprintf("ThrottleMS: %d", *configPtr.ThrottleMs))
	configPtr.v(fmt.Sprintf("Download: %t", *configPtr.Download))
	configPtr.v(fmt.Sprintf("JSON: %t", *configPtr.Json))
	configPtr.v(fmt.Sprintf("Verbose: %t", *configPtr.Verbose))

	// space or command delim
	var bucketNames []string
	if strings.Index(*configPtr.BucketNames, ",") > -1 {
		bucketNames = strings.Split(*configPtr.BucketNames, ",")
	} else {
		bucketNames = strings.Split(*configPtr.BucketNames, " ")
	}

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	var buckets = []*bucketscanner.Bucket{}

	scanners := getScanner(configPtr.CloudProvider)

	for idx, scanner := range scanners {
		wg.Add(idx + 1)

		go func(scanner bucketscanner.Scanner) {
			defer wg.Done()

			for idx, bucketName := range bucketNames {
				bucketName = strings.Trim(bucketName, " ")

				if idx > 0 && *configPtr.ThrottleMs > 0 {
					configPtr.v(fmt.Sprintf("Throttle via sleep for %d MS", *configPtr.ThrottleMs))
					time.Sleep(time.Duration(*configPtr.ThrottleMs) * time.Millisecond)
				}

				configPtr.v(fmt.Sprintf("Getting from %s bucket: %s", scanner.GetProviderName(), bucketName))

				bucket, err := scanner.Get(bucketName)

				if err != nil {
					fmt.Println(err)
				} else {
					configPtr.v("Bucket response received")
					mutex.Lock()
					buckets = append(buckets, bucket)
					mutex.Unlock()
				}
			}
		}(scanner)
	}

	wg.Wait()
	configPtr.v("*** Scan Completed ****")

	// Output Results
	for _, bucket := range buckets {
		if *configPtr.Json {
			jsonStr, err := json.Marshal(bucket)
			if err != nil {
				fmt.Println(err)
				// TODO panic or not?
			}
			fmt.Printf("%s\n", string(jsonStr))
		} else {
			fmt.Printf("%v\n", bucket)
		}
	}

	os.Exit(Success)
}
