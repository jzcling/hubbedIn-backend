.PHONY: plugin

plugin:
	cd krakend && \
    GOOS=linux && \
    GOARCH=amd64 && \
	go build -buildmode=plugin -o grpc-gateway.so .