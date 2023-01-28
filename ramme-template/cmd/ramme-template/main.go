// Package main contains entrypoint app
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/service"
	"github.com/alexadastra/ramme/system"

	impl "github.com/alexadastra/ramme-template/internal/app/service"
	"github.com/alexadastra/ramme-template/internal/swagger"

	"github.com/alexadastra/ramme-template/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Fetch flags configuration
	args := config.ParseFlags()
	config.ServiceName = args.ServiceName
	config.File = args.ConfigPath

	conf, confStart, confStop, err := config.NewConfigFromYAML(args.ConfigPath)
	if err != nil {
		panic(err)
	}

	g := system.NewGroupOperator()
	g.Add(func() error { return confStart(ctx) }, func(err error) { _ = confStop() })

	userGrpcServer := impl.NewRammeTemplate()

	baseGrpcServer := grpc.NewServer()
	api.RegisterRammeTemplateServer(baseGrpcServer, userGrpcServer)

	mux := setupGRPCGateway(ctx, userGrpcServer)
	mux = setupSwagger(mux, conf)

	service.Run(
		ctx,
		g,
		baseGrpcServer,
		mux,
		conf,
	)
}

func setupGRPCGateway(ctx context.Context, userGrpcServer api.RammeTemplateServer) *http.ServeMux {
	mux := http.NewServeMux()
	rmux := runtime.NewServeMux()
	mux.Handle("/", rmux)

	err := api.RegisterRammeTemplateHandlerServer(ctx, rmux, userGrpcServer)
	if err != nil {
		log.Fatal(err)
	}

	return mux
}

func setupSwagger(mux *http.ServeMux, conf config.Config) *http.ServeMux {
	// TODO: this prefix workaround should be solved better
	if conf.Get(config.IsLocalEnvironment).ToBool() {
		mux.Handle(swagger.Pattern, swagger.HandlerLocal)
	} else {
		mux.Handle(swagger.Pattern, swagger.HandlerK8S)
	}

	return mux
}
