#!/bin/bash
set -e

~/go/bin/golangci-lint run

go build -v ./...
go test -v ./...
