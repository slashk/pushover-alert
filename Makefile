BINNAME=pushover-alert

default: build

test:
	go test -v -coverprofile=coverage.out ./...

build:
	go build ./...

release: test
	goreleaser release --rm-dist

docker:
	docker build --tag slashk/pushover-alert:latest .
	docker push  slashk/pushover-alert:latest