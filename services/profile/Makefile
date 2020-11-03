.PHONY: proto
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

profile.pb.go: pb/profile.proto | $(PROTOC_GEN_GO)
	protoc --go_out=plugins=grpc:pb -I pb pb/profile.proto

proto: profile.pb.go