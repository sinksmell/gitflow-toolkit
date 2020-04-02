# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
OUTPUT_DIR := ./bin
TARGET := gitflow-toolkit
HOME := $(HOME)

#
# Define all targets. At least the following commands are required:
#

# All targets.
.PHONY: lint test build  fmt clean

build: build-local

build-local:
	  go build -i -v -o $(OUTPUT_DIR)/$(TARGET)  -ldflags "-s -w " ./

# more info about `GOGC` env: https://github.com/golangci/golangci-lint#memory-usage-of-golangci-lint
lint: $(GOLANGCI_LINT)
	@GOGC=5 $(GOLANGCI_LINT) run

$(GOLANGCI_LINT):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.16.0

fmt:
	goimports -w ./
clean:
	@-rm -vrf ${OUTPUT_DIR}

install: build
	sudo HOME=$(HOME) $(OUTPUT_DIR)/$(TARGET) install

uninstall: build
	sudo HOME=$(HOME) $(OUTPUT_DIR)/$(TARGET) uninstall