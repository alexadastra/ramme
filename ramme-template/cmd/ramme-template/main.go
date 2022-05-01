package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"git.miem.hse.ru/786/ramme-template/internal/app"

	"git.miem.hse.ru/786/ramme/config"
	"git.miem.hse.ru/786/ramme/service"
	"git.miem.hse.ru/786/ramme/system"

	advanced "git.miem.hse.ru/786/ramme-template/internal/config"
	"git.miem.hse.ru/786/ramme-template/internal/swagger"

	"git.miem.hse.ru/786/ramme-template/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// Load ENV configuration
	confManager, confWatcher, err := config.InitBasicConfig()
	if err != nil {
		panic(err)
	}
	cfg := confManager.GetBasic()

	advancedConfManager, advancedConfWatcher, err := advanced.InitAdvancedConfig()
	if err != nil {
		panic(err)
	}
	advancedConfig := advancedConfManager.Get()

	// Configure service and get router
	router, logger, err := service.Setup(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// TODO: figure out how to use advanced config
	logger.Info(advancedConfig)

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

	g.Add(func() error {
		return confWatcher.Run()
	}, func(err error) {
		_ = confWatcher.Close()
	})
	g.Add(func() error {
		return advancedConfWatcher.Run()
	}, func(err error) {
		_ = advancedConfWatcher.Close()
	})

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort))
	if err != nil {
		logger.Fatal(err)
	}
	g.Add(func() error {
		logger.Warnf("Serving grpc address %s", fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort))
		return baseGrpcServer.Serve(grpcListener)
	}, func(error) {
		_ = grpcListener.Close()
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
			_ = httpListener.Close()
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
