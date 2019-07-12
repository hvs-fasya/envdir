build: dep_ensure
	CGO_ENABLED=0 GOOS=linux go build -o go-envdir

GO_TEST_PATHS := $(shell command go list ./... | grep -v "vendor")
test: build
	CGO_ENABLED=0 GOOS=linux go build -o some_prog ./test_prog/test_prog.go
	go test $(GO_TEST_PATHS)

clean:
	rm go-envdir && \
	rm -rf vendor

HAS_DEP := $(shell command -v dep;)
dep_ensure:
ifndef HAS_DEP
	go get -u -v -d github.com/golang/dep/cmd/dep && \
	go install -v github.com/golang/dep/cmd/dep
endif
	dep ensure

.DEFAULT_GOAL := build