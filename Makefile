BINARY_NAME	:= golagen
OUT_FOLDER  := out/bin

build:
	@go build -o $(OUT_FOLDER)/$(BINARY_NAME) cmd/golagen/main.go

clean:
	@go clean
	@rm $(OUT_FOLDER)/$(BINARY_NAME)

run: build
	@$(OUT_FOLDER)/$(BINARY_NAME)

vendor:
	@go mod vendor

.DEFAULT_GOAL = vendor

.PHONY: build \
		clean \
		run \
		vendor
