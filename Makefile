# ########################################################## #
# Makefile for Golang Project
# Includes cross-compiling, installation, cleanup
# ########################################################## #

# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd uname date
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

VERSION=$(shell cat version.txt)
BIN_DIR=bin
BINARY=bucketscanner
COVERAGE_DIR=coverage
COV_PROFILE=${COVERAGE_DIR}/test-coverage.txt
COV_HTML=${COVERAGE_DIR}/test-coverage.html
LIB=lib${BINARY}
PKG=gitlab.com/cjbarker/${BINARY}
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64
BUILD=`date +%FT%T%z`
UNAME=$(shell uname)
GOLIST=$(shell go list ./...)

# Binary Build 
LDFLAGS=-ldflags "-X ${PKG}.Version=${VERSION} -X ${PKG}.Build=${BUILD}"
# Test Build
LDFLAGS_TEST=-ldflags "-X ${PKG}.Version=${VERSION} -X ${PKG}.Build=${BUILD}"


.PHONY: list check clean install build_all all cyclo

default: build

all: clean lint build_all test install

list:
	@echo "Available GNU make targets..."
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

glide:
	# Load Dependencies from glide.lock
	# Run glide up if want to update via glide.yaml
	glide install

vet:
	go vet ${PKG}

lint:
	go get golang.org/x/lint/golint
	go get github.com/GoASTScanner/gas/cmd/gas
	golint -set_exit_status ${GOLIST}
	#gas ${PKG_LOADER}

	golint ${PKG}

format:
	go fmt ${PKG}

build: glide format vet
	go build -o ${BIN_DIR}/${LIB} ${LDFLAGS} ${PKG}
	go build -o ${BIN_DIR}/${BINARY} ${LDFLAGS} ${PKG}/cli

build_all: get format vet
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o $(BIN_DIR)/$(BINARY)-$(GOOS)-$(GOARCH) $(LDFLAGS) $(PKG)/cli; go build -v -o $(BIN_DIR)/$(LIB)-$(GOOS)-$(GOARCH) $(LDFLAGS) $(PKG))))

test: build cyclo
	# tests and code coverage
	mkdir -p $(COVERAGE_DIR)
	go test ${GOLIST} -short -v ${LDFLAGS_TEST} -coverprofile ${COV_PROFILE}
	go tool cover -html=${COV_PROFILE} -o ${COV_HTML}
ifeq ($(UNAME), Darwin)
	open ${COV_HTML}
endif

cyclo:
	@go get github.com/fzipp/gocyclo
	@cyclo_results=$(shell gocyclo -over 20 . | grep -v "vendor")
ifeq ($(cyclo_results),)
	@# ignore no results
else
	printf ${cyclo_results}
endif 

docs: 
	go get golang.org/x/tools/cmd/godoc
	open http://localhost:6060/pkg/github.com/cjbarker/bucketscanner/
	godoc -http=":6060"
	
install:
	go install -o ${LIB} ${LDFLAGS} ${PKG}
	go install -o ${BINARY} ${LDFLAGS} ${PKG}/cli

# Remove only what we've created
clean:
	if [ -d ${BIN_DIR} ] ; then rm -rf ${BIN_DIR} ; fi
	if [ -d ${COVERAGE_DIR} ] ; then rm -rf ${COVERAGE_DIR} ; fi
