.PHONY: proto migration-up migration-down mocks
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

project.pb.go: pb/project.proto | $(PROTOC_GEN_GO)
	protoc -I pb -I $(GOPATH)/src/ --go_out plugins=grpc:../../ pb/project.proto
	protoc -I pb -I $(GOPATH)/src/ --grpc-gateway_out ../../ --grpc-gateway_opt logtostderr=true pb/project.proto
	protoc -I pb -I $(GOPATH)/src/ --openapiv2_out ./pb/ --openapiv2_opt logtostderr=true pb/project.proto

proto: project.pb.go

migration-up:
	cd scripts/migrations && go run . 

migration-down:
	cd scripts/migrations && go run . down

mocks:
	mockery --all && rm -rf tests/mocks && mv mocks tests/