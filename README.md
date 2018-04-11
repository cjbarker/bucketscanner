# Bucket Scanner

[![pipeline status](https://gitlab.com/cjbarker/bucketscanner/badges/master/pipeline.svg)](https://gitlab.com/cjbarker/bucketscanner/commits/master) 
[![coverage report](https://gitlab.com/cjbarker/bucketscanner/badges/master/coverage.svg)] (https://cjbarker.gitlab.io/bucketscanner/test-coverage.html)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/cjbarker/bucketscanner)](https://goreportcard.com/report/gitlab.com/cjbarker/bucketscanner)
[![GitLab license](https://img.shields.io/badge/license-Apache2.0-brightgreen.svg)](https://gitlab.com/cjbarker/bucketscanner/blob/master/LICENSE)

## Overview
*Searching Cloud Storage Since 2017*

## Usage
The bucketscanner requires a cloud provider and action coupled with the bucket name(s).  

```bash
Usage: ./bucketscanner-darwin-amd64 --cloud=CLOUD --action=ACTION [<flags>] <bucket-name>

Cloud command-line bucket (object) scanner.

Flags:
  --help               Show context-sensitive help (also try --help-long and --help-man).
  --version            Show application version.
  --cloud=CLOUD        Cloud provider to scan: aws, gcp, azure. Defaults to all.
  --action=ACTION      Scan action to invoke against bucket: (r)ead, (w)rite, all. Defaults to all.
  --throttle=THROTTLE  Time in milliseconds to throttle subsequent requests sent to a given provider.
  --download           Download bucket content(s).
  --output=OUTPUT      Download bucket content(s) destination directory. Defaults to current user's directory if
                       none passed.
  --JSON               Output results in JSON.
  --verbose            Verbose output messages. Defaults to quiet.

Args:
  <bucket-name>  Bucket(s) name(s) to scan. Does support comma separated for multiple buckets.
```

Example searching one bucket on AWS for read-access:

```bash
./bucketscanner --cloud=aws --action=read --JSON listing-test
{
    "files": [
        {
            "Body": null,
            "directory": true,
            "files": null,
            "name": "empty folder/",
            "size": 0
        },
        {
            "Body": null,
            "directory": true,
            "files": null,
            "name": "empty folder/empty folder/",
            "size": 0
        },
        {
            "Body": null,
            "directory": false,
            "files": null,
            "name": "index-bucketname.html",
            "size": 362
        },
        {
            "Body": null,
            "directory": false,
            "files": null,
            "name": "index-null.html",
            "size": 323
        },
        {
            "Body": null,
            "directory": false,
            "files": null,
            "name": "index-path.html",
            "size": 385
        },
        {
            "Body": null,
            "directory": false,
            "files": null,
            "name": "index-vh.html",
            "size": 385
        }
    ],
    "name": "listing-test",
    "noFiles": 6,
    "provider": "Amazon Simple Storage Service (S3)",
    "scanned": "2018-04-11T11:31:16.78290151-07:00",
    "state": 3,
    "totalSize": 1455,
    "uri": "https://listing-test.s3.amazonaws.com"
}
```

## Developer
Bucketscanner supports multiple platform builds via GNU Make. It does assume and rely on
[Glide](https://github.com/Masterminds/glide) for GoLang package management including dependencies.  Please ensure glide is installed and available in your path before continuing.

To build the binary and library you'll need to clone the repo, setup GoLang and run make.

The default make target builds both components command line binary and library (bucketscanner and libbucketscanner).

```
# Assumes GOPATH exists and golang installed with tools in path
# export PATH=${GOPATH}/bin:${PATH}

cd ${GOPATH}/src
mkdir -p gitlab.com/cjbarker/
cd gitlab.com/cjbarker
git clone git@gitlab.com:cjbarker/bucketscanner.git
cd bucketscanner
make

# Built Binary & Library
ls bin/
bucketscanner libbucketscanner
```

## Continous Integration
[GitLab's CI Pipelines](https://docs.gitlab.com/ee/ci/pipelines.html) handle the continuos integration (CI) and eventually will also handle the continuous deployment (CD) to cloud provider (TBA).

All management of CI/CD is handled via the [.gitlab-ci.yml](https://gitlab.com/cjbarker/bucketscanner/blob/master/.gitlab-ci.yml) file. For more details on  GitLab CI and job configuration consult:  https://docs.gitlab.com/ce/ci/yaml/README.html

Any commit to a branch will trigger the CI.  If you do not want the pipeline's job(s) to trigger you can add a [skip](https://docs.gitlab.com/ee/ci/yaml/README.html#skipping-jobs) to your git commit message.

```bash
git add <file>
git commit -m "[skip ci] will not trigger GitLab CI job"
git push
```
