module github.com/alexadastra/ramme-template

go 1.17

replace (
	github.com/alexadastra/ramme => ../
	github.com/alexadastra/ramme-template/pkg/api => ./pkg/api
)

require (
	github.com/alexadastra/ramme v0.0.0-00010101000000-000000000000
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/pkg/errors v0.9.1
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
)
