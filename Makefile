GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

all: dependencies test build-cli build-service cli-container service-container

dependencies:
	@GOPATH=$(GOPATH) go get -v -t -d ./...

build-cli:
	@GOPATH=$(GOPATH) go build -v -o bin/cli github.com/abrie/mapatlapi/cmd/cli

build-service:
	@GOPATH=$(GOPATH) go build -v -o bin/cli github.com/abrie/mapatlapi/cmd/service

cli-container:
	docker build . -f cli.Dockerfile -t docker.pkg.github.com/abrie/mapatlapi/mapatlcli:latest

service-container:
	docker build . -f service.Dockerfile -t docker.pkg.github.com/abrie/mapatlapi/mapatlservice:latest

test:
	@GOPATH=$(GOPATH) go test -v ./...
