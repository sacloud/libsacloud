TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
export GO111MODULE=on

default: test vet

run:
	go run $(CURDIR)/main.go --disable-healthcheck $(ARGS)

test:
	go test ./sacloud $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-api: 
	go test ./api $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-builder:
	go test ./builder $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-utils: 
	go test ./utils/* $(TESTARGS) -v -timeout=120m -parallel=4 ;

test-all: goimports vet test test-api test-builder test-utils

vet: golint
	go vet ./...

golint: 
	test -z "$$(golint ./... | grep -v 'vendor/' | grep -v '_string.go' | tee /dev/stderr )"

goimports: fmt
	goimports -l -w $(GOFMT_FILES)

fmt:
	gofmt -s -l -w $(GOFMT_FILES)


godoc:
	docker-compose up godoc

.PHONY: default test vet fmt golint test-api test-builder test-all run goimports
