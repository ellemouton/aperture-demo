install:
	go install ./cmd/app
	go install ./cmd/appcli

build:
	go build ./cmd/app
	go build ./cmd/appcli