PROJECT=github.com/alexadastra/ramme

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
PATH_PREFIX?=../

.PHONY: collect
collect:
	mkdir -p ${PATH_PREFIX}${APP}
	cp -r ./ramme-template/. ${PATH_PREFIX}${APP}/
	mv ${PATH_PREFIX}${APP}/cmd/ramme-template ${PATH_PREFIX}${APP}/cmd/${APP}
	mv ${PATH_PREFIX}${APP}/api/ramme-template.proto ${PATH_PREFIX}${APP}/api/${APP}.proto
	mv ${PATH_PREFIX}${APP}/internal/app/service/ramme_template.go ${PATH_PREFIX}${APP}/internal/app/service/$(shell echo ${APP} | tr '-' '_' ).go
	mv ${PATH_PREFIX}${APP}/internal/app/service/ramme_template_test.go ${PATH_PREFIX}${APP}/internal/app/service/$(shell echo ${APP} | tr '-' '_')_test.go
	find ${PATH_PREFIX}${APP} -type f -exec sed -i "s/Ramme Template/${TITLE}/g" {} \;
	find ${PATH_PREFIX}${APP} -type f -exec sed -i "s/RAMME-TEMPLATE/$(shell echo ${APP} | tr '[:lower:]' '[:upper:]')/g" {} \;
	find ${PATH_PREFIX}${APP} -type f -exec sed -i "s/ramme-template/${APP}/g" {} \;
	find ${PATH_PREFIX}${APP} -type f -exec sed -i "s/ramme_template/${SNAKE_CASE_APP}/g" {} \;
	find ${PATH_PREFIX}${APP} -type f -exec sed -i "s/RammeTemplate/${PASCAL_CASE_APP}/g" {} \;
	cd ${PATH_PREFIX}${APP} && make generate APP=${APP}
	cd ${PATH_PREFIX}${APP} && sed -i "s/git.miem.hse.ru\/786\/ramme => ..\//git.miem.hse.ru\/786\/ramme => ..\/ramme/g" go.mod && go mod tidy -compat=1.17


aaa:
	echo ${APP} | tr '-' '_'