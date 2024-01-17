BINARY_NAME	:= golagen
OUT_FOLDER  := ./out/bin
SOURCE_DIR	:= ./cmd/golagen

build:
	@go build -o $(OUT_FOLDER)/$(BINARY_NAME) $(SOURCE_DIR)

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
