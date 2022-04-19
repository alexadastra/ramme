package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"git.miem.hse.ru/786/ramme-template/internal/app"

	"git.miem.hse.ru/786/ramme/config"
	"git.miem.hse.ru/786/ramme/service"
	"git.miem.hse.ru/786/ramme/system"

	"git.miem.hse.ru/786/ramme-template/internal/swagger"

	"git.miem.hse.ru/786/ramme-template/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// Load ENV configuration
	cfg := new(config.Config)
	config.SERVICENAME = "RAMME-TEMPLATE"
	if err := cfg.Load(config.SERVICENAME); err != nil {
		log.Fatal(err)
	}

	// Configure service and get router
	router, logger, err := service.Setup(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// Setup gRPC servers.
	baseGrpcServer := grpc.NewServer()
	userGrpcServer := app.NewRammeTemplate()
	api.RegisterRammeTemplateServiceServer(baseGrpcServer, userGrpcServer)

	// Setup gRPC gateway.
	ctx := context.Background()
	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	{
		err = api.RegisterRammeTemplateServiceHandlerServer(ctx, rmux, userGrpcServer)
		if err != nil {
			logger.Fatal(err)
		}
	}
	mux.Handle(swagger.Pattern, swagger.Handler)

	// Setup secondary HTTP handlers
	// Listen and serve handlers
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.HTTPSecondaryPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Serve
	g := system.NewGroupOperator()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort))
	if err != nil {
		logger.Fatal(err)
	}
	g.Add(func() error {
		logger.Warnf("Serving grpc address %s", fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort))
		return baseGrpcServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

	httpListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.HTTPPort))
	if err != nil {
		logger.Fatal(err)
	}
	g.Add(func() error {
		logger.Warnf("Serving http address %s", fmt.Sprintf("%s:%d", cfg.Host, cfg.HTTPPort))
		return http.Serve(httpListener, mux)
	},
		func(err error) {
			httpListener.Close()
		})

	g.Add(func() error {
		return srv.ListenAndServe()
	}, func(err error) {})

	signals := system.NewSignals()
	g.Add(func() error {
		return signals.Wait(logger, g)
	}, func(error) {})

	if err := g.Run(); err != nil {
		logger.Fatal(err)
	}
}
