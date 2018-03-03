# Bucketscanner development environment within a Docker Container

FROM golang:1.10.0-alpine

LABEL version="1.0.0" \
  maintainer="CJ Barker" \
  maintainer.email="cjbarker@gmail.com" \
  description="Builds Generic Docker image for Bucketscanner (GoLang) Development" \
  repository="https://gitlab.com/cjbarker/bucketscanner"

# Install packages
RUN apk update
RUN apk upgrade
RUN apk add --no-cache --update alpine-sdk

# Install the static analysis tool(s)…
RUN go get -u -v github.com/GoASTScanner/gas
RUN go get -u -v golang.org/x/tools/cmd/goimports
RUN go get -u -v github.com/golang/lint/golint
RUN go get -u -v github.com/alecthomas/gometalinter

# Install Glide for Go Dependency Mgmt
RUN curl https://glide.sh/get | sh

WORKDIR /go/src/gitlab.com/cjbarker/bucketscanner
