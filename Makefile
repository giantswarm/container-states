PROJECT=container-states

BUILD_PATH := $(shell pwd)/.gobuild

GS_PATH := "$(BUILD_PATH)/src/github.com/giantswarm"

BIN=container-states

.PHONY:clean get-deps fmt run-tests

GOPATH := $(BUILD_PATH)

SOURCE=$(shell find . -name '*.go')

all: get-deps $(BIN)

ci: clean all run-tests

clean:
	rm -rf $(BUILD_PATH) $(BIN)

get-deps: .gobuild

.gobuild:
	mkdir -p $(GS_PATH)
	cd "$(GS_PATH)" && ln -s ../../../.. $(PROJECT)
	#
	# Fetch private packages first (so `go get` skips them later)
	#
	# Fetch public dependencies via `go get`

$(BIN): $(SOURCE)
	GOPATH=$(GOPATH) go build -o $(BIN)

run-tests:
	GOPATH=$(GOPATH) go test ./...

fmt:
	gofmt -l -w .
