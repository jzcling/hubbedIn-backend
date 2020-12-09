FROM golang:1.15-alpine as builder
LABEL stage=builder
RUN apk add --no-cache ca-certificates git build-base
WORKDIR /go/src/in-backend/

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY ./helpers/ ./helpers/
COPY ./services/assessment/ ./services/assessment/
RUN cd services/assessment/cmd && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/main

RUN cd services/assessment/scripts/migrations && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/migrations

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /build/dist/ .
COPY --from=builder /go/src/in-backend/services/assessment/configs/config.env .
COPY --from=builder /go/src/in-backend/services/assessment/scripts/migrations/*.sql ./
RUN chmod +x ./main
EXPOSE 50051
ENTRYPOINT ["./main"]