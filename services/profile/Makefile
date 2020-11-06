.PHONY: proto migration-up migration-down
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

profile.pb.go: models/pb/profile.proto | $(PROTOC_GEN_GO)
	protoc --go_out=plugins=grpc:. -I models/pb models/pb/profile.proto

proto: profile.pb.go

migration-up:
	cd scripts/migrations && go run . 

migration-down:
	cd scripts/migrations && go run . down