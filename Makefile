GOPATH=$(shell pwd)
GOBIN=$(shell pwd)/bin

dependencies:
	@GOPATH=$(GOPATH) go get -v -t -d ./...

build:
	@GOPATH=$(GOPATH) go build -v -o bin/cli github.com/abrie/mapatlapi/cmd/cli

container:
	docker build . -t docker.pkg.github.com/abrie/mapatlapi/mapatlcli:latest

test:
	@GOPATH=$(GOPATH) go test -v ./...
