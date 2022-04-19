module git.miem.hse.ru/786/ramme-template

go 1.16

replace (
	git.miem.hse.ru/786/auth-service/pkg/api => ../../auth-service/pkg/api
	git.miem.hse.ru/786/ramme => ../
)

require (
	git.miem.hse.ru/786/auth-service/pkg/api v0.0.0-00010101000000-000000000000
	git.miem.hse.ru/786/ramme v0.0.0-00010101000000-000000000000
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/pkg/errors v0.9.1
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
