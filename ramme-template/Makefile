APP=ramme-template
PROJECT=git.miem.hse.ru/786/ramme-template
REGISTRY?=registry.miem.hse.ru/786/${APP}
CA_DIR?=certs

# Use the 0.0.0 tag for testing, it shouldn't clobber any release builds
RELEASE?=0.0.3
GOOS?=linux
GOARCH?=amd64

RAMME-TEMPLATE_HOST?=0.0.0.0
RAMME-TEMPLATE_HTTP_PORT?=8080
RAMME-TEMPLATE_HTTP_SECONDARY_PORT?=8081
RAMME-TEMPLATE_GRPC_PORT?=6560
RAMME-TEMPLATE_LOG_LEVEL?=0

# Namespace: dev, prod, release, cte, username ...
NAMESPACE?=default

# Infrastructure: dev, stable, test ...
INFRASTRUCTURE?=minikube
VALUES?=values-${INFRASTRUCTURE}

CONTAINER_IMAGE?=${REGISTRY}/${APP}
CONTAINER_NAME?=${APP}-${NAMESPACE}

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

BUILDTAGS=

.PHONY: all
all: build

.PHONY: vendor
vendor: clean bootstrap
	dep ensure

.PHONY: build
build: test certs
	@echo "+ $@"
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/internal/version.RELEASE=${RELEASE} -X ${PROJECT}/internal/version.COMMIT=${COMMIT} -X ${PROJECT}/internal/version.REPO=${REPO_INFO}" \
		-o bin/${GOOS}-${GOARCH}/${APP} ./cmd/${APP}/main.go
	docker build --pull -t $(CONTAINER_IMAGE):$(RELEASE) .

.PHONY: certs
certs:
ifeq ("$(wildcard $(CA_DIR)/ca-certificates.crt)","")
	@echo "+ $@"
	@docker run --name ${CONTAINER_NAME}-certs -d alpine:latest sh -c "apk --update upgrade && apk add ca-certificates && update-ca-certificates"
	@docker wait ${CONTAINER_NAME}-certs
	@mkdir -p ${CA_DIR}
	@docker cp ${CONTAINER_NAME}-certs:/etc/ssl/certs/ca-certificates.crt ${CA_DIR}
	@docker rm -f ${CONTAINER_NAME}-certs
endif

.PHONY: push
push: build
	@echo "+ $@"
	@docker push $(CONTAINER_IMAGE):$(RELEASE)

.PHONY: run
run: build
	@echo "+ $@"
	@docker run --name ${CONTAINER_NAME} -p ${RAMME-TEMPLATE_HTTP_PORT}:${RAMME-TEMPLATE_HTTP_PORT} -p ${RAMME-TEMPLATE_HTTP_SECONDARY_PORT}:${RAMME-TEMPLATE_HTTP_SECONDARY_PORT} -p ${RAMME-TEMPLATE_GRPC_PORT}:${RAMME-TEMPLATE_GRPC_PORT} \
		-e "RAMME-TEMPLATE_HOST=${RAMME-TEMPLATE_HOST}" \
		-e "RAMME-TEMPLATE_HTTP_PORT=${RAMME-TEMPLATE_HTTP_PORT}" \
		-e "RAMME-TEMPLATE_HTTP_SECONDARY_PORT=${RAMME-TEMPLATE_HTTP_SECONDARY_PORT}" \
		-e "RAMME-TEMPLATE_GRPC_PORT=${RAMME-TEMPLATE_GRPC_PORT}" \
		-e "RAMME-TEMPLATE_LOG_LEVEL=${RAMME-TEMPLATE_LOG_LEVEL}" \
		-d $(CONTAINER_IMAGE):$(RELEASE)
	@sleep 1
	@docker logs ${CONTAINER_NAME}

HAS_RUNNED := $(shell docker ps | grep ${CONTAINER_NAME})
HAS_EXITED := $(shell docker ps -a | grep ${CONTAINER_NAME})

.PHONY: logs
logs:
	@echo "+ $@"
	@docker logs ${CONTAINER_NAME}

.PHONY: stop
stop:
ifdef HAS_RUNNED
	@echo "+ $@"
	@docker stop ${CONTAINER_NAME}
endif

.PHONY: start
start: stop
	@echo "+ $@"
	@docker start ${CONTAINER_NAME}

.PHONY: rm
rm:
ifdef HAS_EXITED
	@echo "+ $@"
	@docker rm ${CONTAINER_NAME}
endif

.PHONY: deploy
deploy: push
	helm upgrade ${CONTAINER_NAME} -f charts/${VALUES}.yaml charts --kube-context ${INFRASTRUCTURE} --namespace ${NAMESPACE} --version=${RELEASE} --install --wait --debug

.PHONY: kube_clean
kube_clean:
	kubectl delete -n ${NAMESPACE} deployment ${APP}
	kubectl delete -n ${NAMESPACE} service ${APP}

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

.PHONY: docker_clean
docker_clean: stop rm

.PHONY: generate
generate:
	sudo rm -rf pkg
	sudo rm -rf internal/swagger/*.json
	docker-compose -f docker-compose.yml up
	sudo cp -r ./pkg/api/pkg/api/* ./pkg/api
	sudo rm -rf pkg/api/pkg
	sudo mv ./pkg/api/*.json ./internal/swagger
	sudo chmod -R 777 ./pkg/api
	sudo chmod -R 777 ./internal/swagger
	sed -i 's/v1/${NAMESPACE}\/${APP}\/v1/g' ./internal/swagger/*.json