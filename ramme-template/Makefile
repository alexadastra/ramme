APP=ramme-template
# TODO: add monorepos support (should inclde additional path parts)
PROJECT=github.com/alexadastra/${APP}
REGISTRY?=ghcr.io/alexadastra/${APP}
CA_DIR?=certs

# Use the 0.0.0 tag for testing, it shouldn't clobber any release builds
RELEASE?=0.0.3
GOOS?=linux
GOARCH?=amd64

# Namespace: dev, prod, release, cte, username ...
NAMESPACE?=default

# Infrastructure: dev, stable, test ...
INFRASTRUCTURE?=minikube
VALUES?=values-${INFRASTRUCTURE}

CONTAINER_IMAGE?=${REGISTRY}/${APP}
CONTAINER_NAME?=${APP}-${NAMESPACE}

HTTP_PORT?=8080
HTTP_SECONDARY_PORT?=8081
GRPC_PORT?=8082

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

MANAGER?=docker

BUILDTAGS=

.PHONY: all
all: build

.PHONY: build
build: test certs
	@echo "+ $@"
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/internal/version.RELEASE=${RELEASE} -X ${PROJECT}/internal/version.COMMIT=${COMMIT} -X ${PROJECT}/internal/version.REPO=${REPO_INFO}" \
		-o bin/${GOOS}-${GOARCH}/${APP} ./cmd/${APP}/main.go
	${MANAGER} build --pull -t $(CONTAINER_IMAGE):$(RELEASE) .

.PHONY: certs
certs:
ifeq ("$(wildcard $(CA_DIR)/ca-certificates.crt)","")
	@echo "+ $@"
	@${MANAGER} run --name ${CONTAINER_NAME}-certs -d alpine:latest sh -c "apk --update upgrade && apk add ca-certificates && update-ca-certificates"
	@${MANAGER} wait ${CONTAINER_NAME}-certs
	@mkdir -p ${CA_DIR}
	@${MANAGER} cp ${CONTAINER_NAME}-certs:/etc/ssl/certs/ca-certificates.crt ${CA_DIR}
	@${MANAGER} rm -f ${CONTAINER_NAME}-certs
endif

.PHONY: push
push: build
	@echo "+ $@"
	@${MANAGER} push $(CONTAINER_IMAGE):$(RELEASE)

.PHONY: run
run: build
	@echo "+ $@"
	@${MANAGER} run --name ${CONTAINER_NAME} -p ${HTTP_PORT}:${HTTP_PORT} \
		-p ${HTTP_SECONDARY_PORT}:${HTTP_SECONDARY_PORT} \
		-p ${GRPC_PORT}:${GRPC_PORT} \
		-d $(CONTAINER_IMAGE):$(RELEASE)
	@sleep 1
	@${MANAGER} logs ${CONTAINER_NAME}

HAS_RUNNED := $(shell ${MANAGER} ps | grep ${CONTAINER_NAME})
HAS_EXITED := $(shell ${MANAGER} ps -a | grep ${CONTAINER_NAME})

.PHONY: logs
logs:
	@echo "+ $@"
	@${MANAGER} logs ${CONTAINER_NAME}

.PHONY: stop
stop:
ifdef HAS_RUNNED
	@echo "+ $@"
	@${MANAGER} stop ${CONTAINER_NAME}
endif

.PHONY: start
start: stop
	@echo "+ $@"
	@${MANAGER} start ${CONTAINER_NAME}

.PHONY: rm
rm:
ifdef HAS_EXITED
	@echo "+ $@"
	@${MANAGER} rm ${CONTAINER_NAME}
endif

.PHONY: deploy
deploy: push
	@echo "+ $@"
	helm upgrade ${CONTAINER_NAME} -f charts/${VALUES}.yaml charts --kube-context ${INFRASTRUCTURE} --namespace ${NAMESPACE} --version=${RELEASE} --install --wait --debug

.PHONY: delete_service
delete_service:
ifdef $(shell kubectl get services | grep ${APP})
	@echo "+ $@"
	kubectl delete -n ${NAMESPACE} service ${APP}
endif

.PHONY: delete_deployment
delete_deployment:
ifdef $(shell kubectl get deployments | grep ${APP})
	@echo "+ $@"
	kubectl delete -n ${NAMESPACE} deployment ${APP}
endif

.PHONY: delete_configmap
delete_configmap:
ifdef $(shell kubectl get deployments | grep ${APP})
	@echo "+ $@"
	kubectl delete -n ${NAMESPACE} configmap ${APP}-configmap
endif

.PHONY: kube_clean
kube_clean: delete_service delete_deployment delete_configmap

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

.PHONY: clean
clean: stop rm
	@rm -f bin/${GOOS}-${GOARCH}/${APP}

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

.PHONY: local_clean
local_clean: stop rm

.PHONY: generate
generate:
	@echo "+ $@"
	@rm -rf pkg
	@rm -rf internal/swagger/*.json
	@${MANAGER} run -v ./:/defs --name protoc-all -d namely/protoc-all:latest \
                -d api -o tmp -i third_party/googleapis -i pkg/api -l go --with-gateway
	@${MANAGER} wait protoc-all
	@${MANAGER} logs protoc-all
	@${MANAGER} rm -f protoc-all
	@mv ./tmp/pkg .
	@mv ./tmp/${APP}.swagger.json ./internal/swagger
	@rm -rf tmp
	@cd ./internal/swagger && cp ./${APP}.swagger.json ./${APP}.swagger.local.json
	@cd ./internal/swagger && mv ./${APP}.swagger.json ./${APP}.swagger.k8s.json
	@cd ./internal/swagger && sed -i '' -e 's/v1/${NAMESPACE}\/${APP}\/v1/g' ./${APP}.swagger.k8s.json
	@cd ./pkg/api && go mod init ${PROJECT}/pkg/api && go mod tidy