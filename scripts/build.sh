#!/bin/bash -e

ORG_PATH="github.com/andersjanmyr"
REPO_PATH="${ORG_PATH}/ocr"

export GOPATH=${PWD}/gopath

rm -f $GOPATH/src/${REPO_PATH}
mkdir -p $GOPATH/src/${ORG_PATH}
ln -s ${PWD} $GOPATH/src/${REPO_PATH}

eval $(go env)
go get \
	cloud.google.com/go/vision/apiv1 \
	github.com/aws/aws-sdk-go/aws \
	github.com/aws/aws-sdk-go/aws/endpoints \
	github.com/aws/aws-sdk-go/aws/session \
	github.com/aws/aws-sdk-go/service/rekognition

CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o bin/ocr ${REPO_PATH}
