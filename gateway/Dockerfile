FROM golang:1.15-buster as builder
LABEL stage=builder
WORKDIR /go/src/in-backend/

# restore dependencies
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY ./gateway/ ./gateway/
COPY ./services/profile/pb ./services/profile/pb
RUN cd gateway/krakend && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -buildmode=plugin -o /build/dist/grpc-gateway.so

FROM devopsfaith/krakend:latest

WORKDIR /etc/krakend

COPY --from=builder /build/dist/ .
COPY --from=builder /go/src/in-backend/gateway/krakend/krakend.json .

ENTRYPOINT [ "/usr/bin/krakend" ]
CMD [ "run", "-c", "/etc/krakend/krakend.json" ]

EXPOSE 8000 8090