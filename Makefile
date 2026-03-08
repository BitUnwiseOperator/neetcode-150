.PHONY: fmt vet test all 

all: fmt vet lint test

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test -v -race ./...
