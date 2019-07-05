TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
export GO111MODULE=on

default: clean gen fmt goimports golint vet test

test:
	TESTACC= go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 ;

testacc:
	TESTACC=1 go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 ;

.PHONY: clean
clean:
	rm -f sacloud/zz_*.go; \
	rm -f sacloud/fake/zz_*.go \
	rm -f sacloud/naked/zz_*.go \
	rm -f sacloud/stub/zz_*.go \
	rm -f sacloud/trace/zz_*.go

.PHONY: gen
gen: clean
	go generate ./...; gofmt -s -l -w $(GOFMT_FILES); goimports -l -w $(GOFMT_FILES)

vet: golint
	go vet ./...

golint: 
	test -z "$$(golint ./... | grep -v 'tools/' | grep -v 'vendor/' | grep -v '_string.go' | tee /dev/stderr )"

goimports: fmt
	goimports -l -w $(GOFMT_FILES)

fmt:
	gofmt -s -l -w $(GOFMT_FILES)

godoc:
	@echo "URL: http://localhost:6060/pkg/github.com/sacloud/libsacloud/"; \
	docker run -it --rm -v $$PWD:/go/src/github.com/sacloud/libsacloud -p 6060:6060 golang:1.12 godoc -http=:6060

.PHONY: tools
tools:
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get golang.org/x/lint/golint
	GO111MODULE=off go get github.com/motemen/gobump

.PHONY: bump-patch bump-minor bump-major version
bump-patch:
	@gobump patch -w ; echo "next version is v`gobump show -r`"

bump-minor:
	@gobump minor -w ; echo "next version is v`gobump show -r`"

bump-major:
	@gobump major -w ; echo "next version is v`gobump show -r`"

version:
	@gobump show -r

git-tag:
	git tag v`gobump show -r`
