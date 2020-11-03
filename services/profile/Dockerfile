FROM golang:1.15-alpine as builder
RUN apk add --no-cache ca-certificates git build-base
WORKDIR /build

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -a -o /build/main .

FROM alpine as release
RUN apk add --no-cache ca-certificates
COPY --from=builder /build/main .
RUN chmod +x ./main
ENTRYPOINT ["./main"]