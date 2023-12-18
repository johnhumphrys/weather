build:
	mkdir -p build
	go build ./...

lint:
	golangci-lint run ./... --fix --config ".golangci.yml" --verbose

install:
	GOBIN=`pwd`/build go install ./...
