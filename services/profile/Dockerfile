FROM golang:1.15-alpine as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /build

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY . /usr/local/go/src/in-backend/services/profile
RUN go build -o /build/main .

FROM alpine as release
RUN apk add --no-cache ca-certificates
COPY --from=builder /build/main /profile
RUN chmod +x /profile
ENTRYPOINT ["/profile"]