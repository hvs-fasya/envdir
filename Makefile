build:
	CGO_ENABLED=0 go build -o go-envdir

GO_TEST_PATHS := $(shell command go list ./... | grep -v "vendor")
test: build
	CGO_ENABLED=0 go build -o some_prog ./test_prog/test_prog.go
	go test $(GO_TEST_PATHS)

clean:
	rm go-envdir && \
	rm -rf vendor

.DEFAULT_GOAL := build