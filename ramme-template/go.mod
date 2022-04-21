module git.miem.hse.ru/786/ramme-template

go 1.16

replace (
	git.miem.hse.ru/786/auth-service/pkg/api => ../../auth-service/pkg/api
	git.miem.hse.ru/786/ramme => ../
)

require (
	git.miem.hse.ru/786/ramme v0.0.0-00010101000000-000000000000
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
