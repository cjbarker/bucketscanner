# Bucket Scanner

[![Build Status](https://travis-ci.org/cjbarker/bucketscanner.svg?branch=master)](https://travis-ci.org/cjbarker/bucketscanner)

----

## Overview
*Searching Cloud Storage Since 2017*

## Developer
Bucket Scanner supports multiple platform builds via GNU Make. 

To build clone the repo, setup Go and run make.

The default make target builds the library and command line binary in the same directory: bucketscanner and libbucketscanner.

```
git clone git@github.com:cjbarker/bucketscanner.git
cd bucketscanner
export GOPATH=${GOPATH}:`pwd`
make
```
