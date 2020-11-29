FROM golang:1.15-alpine as builder
LABEL stage=builder
RUN apk add --no-cache ca-certificates git build-base
WORKDIR /go/src/in-backend/

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY ./helpers/ ./helpers/
COPY ./services/joblisting/ ./services/joblisting/
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

FROM openjdk:alpine as release
RUN apk add --no-cache ca-certificates git curl sed unzip nodejs npm
WORKDIR /app

ARG RELEASE=4.5.0.2216

RUN set -x &&\
  curl --insecure -o ./sonarscanner.zip -L https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-$RELEASE-linux.zip &&\
  unzip sonarscanner.zip &&\
  rm sonarscanner.zip &&\
  rm sonar-scanner-$RELEASE-linux/jre -rf &&\
#   ensure Sonar uses the provided Java for musl instead of a borked glibc one
  sed -i 's/use_embedded_jre=true/use_embedded_jre=false/g' /app/sonar-scanner-$RELEASE-linux/bin/sonar-scanner

ENV SONAR_RUNNER_HOME=/app/sonar-scanner-$RELEASE-linux
ENV SONAR_HOST_URL=http://sonarqube:9000
ENV PATH $PATH:/app/sonar-scanner-$RELEASE-linux/bin

COPY --from=builder /build/dist/ .
COPY --from=builder /go/src/in-backend/services/joblisting/configs/config.env .
COPY --from=builder /go/src/in-backend/services/joblisting/scripts/migrations/*.sql ./
COPY --from=builder /go/src/in-backend/services/joblisting/service/*.sh ./
RUN chmod +x ./main
RUN chmod +x ./*.sh
EXPOSE 50052
ENTRYPOINT ["./main"]