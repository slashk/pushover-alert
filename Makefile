BINNAME=pushover-alert

default: build

test:
	go test -v -coverprofile=coverage.out ./...

build:
	go build ./...

release: test
	goreleaser release --rm-dist
