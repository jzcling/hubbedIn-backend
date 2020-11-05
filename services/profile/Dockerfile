FROM golang:1.15-alpine as builder
LABEL stage=builder
RUN apk add --no-cache ca-certificates git build-base
WORKDIR /build

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN cd cmd && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/main

RUN cd scripts/migrations && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/migrations

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /build/dist/ .
COPY --from=builder /build/configs/env.yaml .
RUN chmod +x ./main
ENTRYPOINT ["./main"]