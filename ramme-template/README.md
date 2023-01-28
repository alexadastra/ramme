# Ramme Template

## Usage


## Makefile variables

- `APP` - actual service name, used everywhere
- `PROJECT` - project repo on Github, Gitlab etc.
- `REGISTRY` - container registry the image will be pushed to;
- `RELEASE` - SemVer version of the release;
- `NAMESPACE` - the namespace where service will be deployed to;
- `INFRASTRUCTURE` - name of config to use while deploying (`charts/values-($INFRASTRUCTURE).yaml` will be used);
- `HTTP_PORT`- preferred port for recieving HTTP requests;
- `HTTP_ADMIN_PORT` - preferred port for internal HTTP requests (docs, health- and readyness check etc.);
- `GRPC_PORT` - preferred port for recieving gRPC requests;

## Makefile commands

### Generating
- `generate` - update generated `pkg/*` files and Swagger docs by current `api/*.proto` files

### Linting & testing
- `bootstrap` - install godep, golint and golang-ci linter on your $GOPATH;
- `test` - run `go fmt`, `go lint`, `go vet`, `go test` on all the Go files in directory;
- `coverage` - count up test coverage;
- `lint_full` - run golang-ci linter;

### Building & running locally
- `certs` - create sertificates;
- `build` - run `test`, `certs`, build the broject into binary and put binary and certs into Docker image;
- `run` - run `build` and run container on local Docker;
- `logs` - fetch logs from inside the container
- `stop` - stop container if running;
- `start` - restart container;
- `rm` - remove container image from Docker;

### Deploying
- `push` - run `build` and push image into registry;
- `deploy` - run `push` and deploy service on Kubernetes by upgrading Helm release with charts;
- `kube_clean` - removes the service, deployment and config from Kubernetes;
