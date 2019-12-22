TEST?=./...

default: test

test-all: test testrace testbench

test: 
	go list $(TEST) | xargs -n1 go test -timeout=60s -parallel=10 $(TESTARGS)

testrace:
	go list $(TEST) | xargs -n1 go test -race $(TESTARGS)

testbench:
	go list $(TEST) | xargs -n1 go test -bench=.

# updatedeps installs all the dependencies to run and build
updatedeps:
	go mod vendor

deps:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run ./...

build:
	cd cmd
	go build -o rd2wgs84

.PHONY: updatedeps test testrace build

