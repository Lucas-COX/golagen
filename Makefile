BINARY_NAME		:= golagen
PACKAGE_NAME 	:= Lucas-COX/golagen
OUT_FOLDER  	:= ./out/bin
SOURCE_DIR		:= ./cmd/golagen
BUILD_TIME		:= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION			?= $(shell git branch --show-current)

build:
	@go build -o $(OUT_FOLDER)/$(BINARY_NAME) -ldflags "  \
	-X $(PACKAGE_NAME)/internal.version=$(VERSION)		 \
	-X $(PACKAGE_NAME)/internal.buildTime=$(BUILD_TIME)" \
	$(SOURCE_DIR)

clean:
	@go clean
	@rm $(OUT_FOLDER)/$(BINARY_NAME)

install:
	@go mod tidy

run: build
	@$(OUT_FOLDER)/$(BINARY_NAME)

test: build
	@cd tests && ../$(OUT_FOLDER)/$(BINARY_NAME) --verbose

vendor:
	@go mod vendor

.DEFAULT_GOAL = install

.PHONY: build \
		clean \
		run \
		vendor
