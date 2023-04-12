.PHONY: generate generate-oapi-codegen generate-all build test

build:
	go build .

generate:
	go generate ./...

generate-all: generate generate-oapi-codegen

generate-oapi-codegen:
	cd examples/oapi-codegen-client && go generate ./...

test:
	go test ./...
