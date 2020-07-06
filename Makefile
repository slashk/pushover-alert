BINNAME=pushover-alert

default: build

test:
	go test -v -coverprofile=coverage.out ./...

release: test
	goreleaser release --rm-dist