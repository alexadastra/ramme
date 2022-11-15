module github.com/alexadastra/ramme-template

go 1.18

replace (
	github.com/alexadastra/ramme => ../
	github.com/alexadastra/ramme-template/pkg/api => ./pkg/api
)

require (
	github.com/alexadastra/ramme v0.0.0-00010101000000-000000000000
	github.com/alexadastra/ramme-template/pkg/api v0.0.0-00010101000000-000000000000
	github.com/flowchartsman/swaggerui v0.0.0-20221017034628-909ed4f3701b
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.50.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20221117204609-8f9c96812029 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
)
