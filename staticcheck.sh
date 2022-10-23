#!/usr/bin/env bash

echo "==> Checking that code complies with static analysis requirements..."

PACKAGES=$(go list ./...)

staticcheck -checks 'all,-ST*' ${PACKAGES} && go vet ./...

golangci-lint run ./...
