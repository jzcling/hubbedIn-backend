FROM golang:1.15-alpine as builder
LABEL stage=builder
RUN apk add --no-cache ca-certificates git build-base
WORKDIR /go/src/in-backend/

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY ./helpers/ ./helpers/
COPY ./services/joblisting/ ./services/joblisting/
COPY ./services/profile/pb ./services/profile/pb
RUN cd services/joblisting/cmd && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/main

RUN cd services/joblisting/scripts/migrations && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/migrations

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /build/dist/ .
COPY --from=builder /go/src/in-backend/services/joblisting/configs/config.env .
COPY --from=builder /go/src/in-backend/services/joblisting/scripts/migrations/*.sql ./
RUN chmod +x ./main
EXPOSE 50051
ENTRYPOINT ["./main"]