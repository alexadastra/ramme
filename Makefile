PROJECT=git.miem.hse.ru/786/ramme

BUILDTAGS=

GO_LIST_FILES=$(shell go list ${PROJECT}/...)

.PHONY: fmt
fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: lint
lint: bootstrap
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint -min_confidence=0.85 {{.Dir}}/..."{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: lint-full
lint-full: lint
	@echo "+ $@"
	@golangci-lint run

.PHONY: vet
vet:
	@echo "+ $@"
	@go vet ${GO_LIST_FILES}

.PHONY: test
test: fmt lint vet
	@echo "+ $@"
	@go test -v -race -cover -tags "$(BUILDTAGS) cgo" ${GO_LIST_FILES}

.PHONY: cover
cover:
	@echo "+ $@"
	@> coverage.txt
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}} && cat {{.Dir}}/.coverprofile  >> coverage.txt"{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c


HAS_DEP := $(shell command -v dep;)
HAS_LINT := $(shell command -v golint;)
HAS_LINT_FULL := $(shell command -v golangci-lint;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_DEP
	go get -u github.com/golang/dep/cmd/dep
endif
ifndef HAS_LINT
	go get -u golang.org/x/lint/golint
endif
ifndef HAS_LINT_FULL
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.44.0
endif

TITLE?="Some Cool App"
APP?=$(shell echo ${TITLE} | tr ' ' '-' | tr '[:upper:]' '[:lower:]')
SNAKE_CASE_APP?=$(shell echo ${APP} | tr '-' '_')
PASCAL_CASE_APP?=$(shell echo ${SNAKE_CASE_APP} | sed -r 's/(^|_)([a-z])/\U\2/g')

.PHONY: collect
collect:
	cp -r ./ramme-template ./${APP}
	mv ./${APP}/cmd/ramme-template ./${APP}/cmd/${APP}
	mv ./${APP}/api/ramme_template.proto ./${APP}/api/${APP}.proto
	mv ./${APP}/internal/app/ramme_template.go ./${APP}/internal/app/$(shell echo ${APP} | tr '-' '_' ).go
	mv ./${APP}/internal/app/ramme_template_test.go ./${APP}/internal/app/$(shell echo ${APP} | tr '-' '_')_test.go
	find ./${APP} -type f -exec sed -i "s/Ramme Template/${TITLE}/g" {} \;
	find ./${APP} -type f -exec sed -i "s/ramme-template/${APP}/g" {} \;
	find ./${APP} -type f -exec sed -i "s/ramme_template/${SNAKE_CASE_APP}/g" {} \;
	find ./${APP} -type f -exec sed -i "s/RammeTemplate/${PASCAL_CASE_APP}/g" {} \;
	cd ./${APP} && make generate APP=${APP}


aaa:
	echo ${APP} | tr '-' '_'