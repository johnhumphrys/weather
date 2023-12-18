FROM golang:1.21.5-bullseye AS builder
# RUN yes | apt-get update && yes | apt-get install awscli ca-certificates curl tzdata && rm -rf /var/cache/apk/*
ENV CGO_ENABLED 0

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . ./

# Builds the app
RUN make build

# generates binary
RUN GOBIN=`pwd`/build go install ./...

FROM debian:12.2 AS weather
RUN yes | apt-get update && yes | apt-get install awscli ca-certificates curl tzdata && rm -rf /var/cache/apk/*
WORKDIR /app
ENV WORKDIR=/app
COPY --from=builder app/build/ .

ARG SERVICE_NAME
ARG DEPLOY_REGION
ARG REVISION
ARG IOT_PLATFORM_URL

HEALTHCHECK CMD curl --fail http://localhost:3000/health || exit 1
# Tells Docker which network port your stubbed listens on
EXPOSE 3000

ENTRYPOINT ["./weather"]