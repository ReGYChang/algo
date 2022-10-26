#!/usr/bin/env bash

echo "==> Checking that code complies with go fmt requirements..."

GOFILES=$(find . -name "*.go" -type f | grep -v "/vendor/")

gofmt -s -l -w ${GOFILES}
