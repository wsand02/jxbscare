all: build build-cli build-web

build:
	go build ./...

build-cli:
	go build -o ./bin/jxbscare-cli ./cmd/jxbscare-cli 

build-web:
	go build -o ./bin/jxbscare-web ./cmd/jxbscare-web 