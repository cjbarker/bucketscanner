# Bucket Scanner

[![pipeline status](https://gitlab.com/cjbarker/bucketscanner/badges/master/pipeline.svg)](https://gitlab.com/cjbarker/bucketscanner/commits/master) 
[![coverage report](https://gitlab.com/cjbarker/bucketscanner/badges/master/coverage.svg)] (https://cjbarker.gitlab.io/bucketscanner/test-coverage.html)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/cjbarker/bucketscanner)](https://goreportcard.com/report/gitlab.com/cjbarker/bucketscanner)
[![GitLab license](https://img.shields.io/badge/license-Apache2.0-brightgreen.svg)](https://gitlab.com/cjbarker/bucketscanner/blob/master/LICENSE)

## Overview
*Searching Cloud Storage Since 2017*

## Usage
TBA

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
