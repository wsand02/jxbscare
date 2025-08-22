all: build build-cli

build:
	go build ./...

build-cli:
	go build -o ./bin/jxbscare-cli ./cmd/jxbscare-cli 