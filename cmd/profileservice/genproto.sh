#!/bin/bash -e

PATH=$PATH:$GOPATH/bin
protodir=../../api/proto/v1

protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/in.proto