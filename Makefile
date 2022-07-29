APPNAME:=$(shell basename $(shell go list))
APP_VERSION:=v0.0.2

GO_LDFLAGS+=-s -w -X 'main.buildVersion=$(APP_VERSION)'
GO_LDFLAGS:=-ldflags="$(GO_LDFLAGS)"
BINARY = ./$(APPNAME)
GO_FILES = $(shell find . -type f -name '*.go')

# Rebuild BINARY only if GO_FILES have been modified
#GOOS=linux GOARCH=amd64 
$(BINARY): $(GO_FILES)
	go build $(GO_LDFLAGS) -o $@

.PHONY: build
build: $(BINARY)

.PHONY: format
format:
	gofmt -s -w .
