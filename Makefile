default: deps test lint

deps:
	go get github.com/alecthomas/gometalinter
	go get github.com/tools/godep
	godep restore

lint:
	gometalinter --install
	gometalinter ./...

test:
	godep go test ./...

.PHONY: default deps lint test
