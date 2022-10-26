GO ?= go
PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)

.PHONY: check fmt

test: fmt
	@$(GO) test -race -cover $(PACKAGES)

fmt:
	@sh -c "'$(CURDIR)/script/gofmt.sh'"

