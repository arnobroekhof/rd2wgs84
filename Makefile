TEST?=./...

default: test

test:
	go list $(TEST) | xargs -n1 go test -timeout=60s -parallel=10 $(TESTARGS)

testrace:
	go list $(TEST) | xargs -n1 go test -race $(TESTARGS)

# updatedeps installs all the dependencies to run and build
updatedeps:
	go mod vendor

build:
	cd cmd
	go build -o rd2wgs84

.PHONY: updatedeps test testrace build
