GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

all: dependencies test build container

dependencies:
	@GOPATH=$(GOPATH) go get -v -t -d ./...

build:
	@GOPATH=$(GOPATH) go build -v -o bin/cli github.com/abrie/mapatlapi/cmd/cli

container:
	docker build . -t abriedev/mapatlapi:latest

test:
	@GOPATH=$(GOPATH) go test -v ./...
