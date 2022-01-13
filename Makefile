APPNAME:=$(shell basename $(shell go list))

GO_LDFLAGS+=-s -w
GO_LDFLAGS:=-ldflags="$(GO_LDFLAGS)"

BINARY = ./$(APPNAME)
GO_FILES = $(shell find . -type f -name '*.go')

# Rebuild BINARY only if GO_FILES have been modified
$(BINARY): $(GO_FILES)
	GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o $@

.PHONY: build
build: $(BINARY)

.PHONY: format
format:
	gofmt -s -w .
