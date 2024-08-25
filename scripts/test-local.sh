#!/bin/bash
set -e

/snap/bin/golangci-lint run

go build -v ./...
go test -v ./...
