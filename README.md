# Ramme
Ramme is an SDK for creating ready-for-production backend microservices. These services are best suitable to be used in Kubernetes and contains all necessary templates for that.

## Tech Stack  

Golang, Docker, Makefile, Helm, Kubernetes

## Features  

- Creates or deploys service with one Makefile command;
- Created service contains:
    - gRPC, HTTP handlers and Swagger page created from single .proto-file;
    - logging;
    - configuration (from file, env or k8s configmap);
    - signal monitoring & graceful shutdown;
    - health- and readyness- probes.
- Provided Makefile commands are:
    - bootstrapping all the necessary go libs into local machine;
    - linting;
    - test running;
    - building;
    - passing certs;
    - creating & publishing Docker container;
    - creating and publishing Helm deployment.

## Usage

### Requirements
In order to operate SDK properly, you'll need:
- A machine with Makefile & Git support;
- Golang v1.17+;
- Container manager (Docker or Podman);
- Helm;
- kubectl with access to the preferrable cluster; 
Docker is set by default to run processes and build an app container. To set it to podman, run:
~~~bash
export MANAGER='podman'
~~~

### Creating service

Clone repository and locate to the repo folder:
~~~bash
git clone https://github.com/alexadastra/ramme && cd ramme
~~~
Set environment variables required to run the command correctly:
~~~bash
export APP='your cool app name'
export PATH_PREFIX='~/path/to/service'
~~~
And just like that run the command:
~~~bash
make collect
~~~
The service is now stored in `PATH_PREFIX/APP` folder

### Generating handlers

Locate to service folder and edit api/my-cool-app.proto

~~~bash
cd my-cool-app && vim api/my-cool-app.proto
~~~
Create rpc`s according to [docs](https://developers.google.com/protocol-buffers/docs/proto3):
~~~Proto
service RammeTemplate {
  rpc Ping(PingRequest) returns (PingResponse) {
    option (google.api.http) = {
      post: "/v1/ping"
      body: "*"
    };
  }
}

message PingRequest {
  string value = 1;
}

message PingResponse {
  int64 code = 1;
  string value = 2;
}
~~~
run the command:
~~~bash
make generate
~~~
And - TAADAAAH! - all the grpc, grpc-gateway middleware is at `pkg/api` folder, and OpenAPI json is at `internal/swagger`.

## Feedback  

If you have any feedback, please reach out to me at its.aleksey.semenchenko@gamil.com

## License  

[MIT](https://choosealicense.com/licenses/mit/)
