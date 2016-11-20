all: build

build:
	@go build -o $(GOPATH)/bin/gobot

test: 
	@go test -v