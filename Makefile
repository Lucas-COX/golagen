BINARY_NAME	:= golagen
OUT_FOLDER  := ./out/bin
SOURCE_DIR	:= ./cmd/golagen
BUILD_TIME	:= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION		:= $(shell git branch --show-current)

build:
	@go build -o $(OUT_FOLDER)/$(BINARY_NAME) -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" $(SOURCE_DIR)

clean:
	@go clean
	@rm $(OUT_FOLDER)/$(BINARY_NAME)

install:
	@go mod tidy

run: build
	@$(OUT_FOLDER)/$(BINARY_NAME)

vendor:
	@go mod vendor

.DEFAULT_GOAL = install

.PHONY: build \
		clean \
		run \
		vendor
