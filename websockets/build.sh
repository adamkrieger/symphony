#!/usr/bin/env bash

# Build executable

GOOS=linux GOARCH=amd64 go build -o ./bin/confcall main.go
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi

# Build Docker image

docker build --no-cache -t confcall:latest .