GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=auth_connector
BINARY_UNIX=$(BINARY_NAME)_unix
GOMOD=$(GOCMD) mod

all: run

build:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)
	rm -f ./build/$(BINARY_UNIX)

run:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./
	./build/$(BINARY_NAME)

restart:
	kill -INT $$(cat pid)
	$(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./
	./build/$(BINARY_NAME)

deps:
	$(GOMOD) download
	$(GOMOD) vendor

cross:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -tags=jsoniter -v ./

docker:
	./docker_build.sh && ./docker_run.sh