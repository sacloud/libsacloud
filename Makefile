TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
export GO111MODULE=on

default: fmt goimports golint vet test

test:
	go test ./sacloud $(TESTARGS) -v -timeout=120m -parallel=4 ;

vet: golint
	go vet ./...

golint: 
	test -z "$$(golint ./... | grep -v 'vendor/' | grep -v '_string.go' | tee /dev/stderr )"

goimports: fmt
	goimports -l -w $(GOFMT_FILES)

fmt:
	gofmt -s -l -w $(GOFMT_FILES)

godoc:
	@echo "URL: http://localhost:6060/pkg/github.com/sacloud/libsacloud-v2/"; \
	docker run -it --rm -v $$PWD:/go/src/github.com/sacloud/libsacloud-v2 -p 6060:6060 golang:1.12 godoc -http=:6060

.PHONY: tools
tools:
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get golang.org/x/lint/golint
	GO111MODULE=off go get github.com/motemen/gobump

.PHONY: bump-patch bump-minor bump-major version
bump-patch:
	gobump patch -w

bump-minor:
	gobump minor -w

bump-major:
	gobump major -w

version:
	gobump show

git-tag:
	git tag v`gobump show -r`
