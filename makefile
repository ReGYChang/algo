GO ?= go
PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)

.PHONY: check fmt

test: check fmt
	@$(GO) test -race -cover $(PACKAGES)

check:
	@sh -c "'$(CURDIR)/script/staticcheck.sh'"

fmt:
	@sh -c "'$(CURDIR)/script/gofmt.sh'"

