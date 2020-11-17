FROM golang:1.15-alpine as builder
LABEL stage=builder
WORKDIR /go/src/in-backend/

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY ./gateway/ ./gateway/
COPY ./services/profile/pb ./services/profile/pb
RUN cd gateway/cmd && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/main

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /build/dist/ .
COPY --from=builder /go/src/in-backend/gateway/configs/config.env .
RUN chmod +x ./main
EXPOSE 8000
ENTRYPOINT ["./main"]