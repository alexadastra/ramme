package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/system"
	"google.golang.org/grpc"
)

// Run inits and starts up ramme app components
func Run(ctx context.Context, g *system.GroupOperator, baseGrpcServer *grpc.Server, mux http.Handler, basicConfig *config.BasicConfig) {
	// Configure service and get router
	router, logger, err := Setup(basicConfig)
	if err != nil {
		log.Fatal(err)
	}

	grpcStart, grpcStop := setupGRPC(baseGrpcServer, basicConfig)
	g.Add(grpcStart, grpcStop)
	logger.Warnf("Serving grpc address %s", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.GRPCPort))

	httpStart, httpStop := setupHTTP(mux, &HttpServerConfig{
		WriteTimeOut: 15 * time.Second,
		ReadTimeOut:  15 * time.Second,
		Host:         basicConfig.Host,
		Port:         basicConfig.HTTPPort,
	})
	g.Add(httpStart, httpStop)
	logger.Warnf("Serving http address %s", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.HTTPPort))

	httpSecStart, httpSecStop := setupHTTP(router, &HttpServerConfig{
		WriteTimeOut: 15 * time.Second,
		ReadTimeOut:  15 * time.Second,
		Host:         basicConfig.Host,
		Port:         basicConfig.HTTPSecondaryPort,
	})
	g.Add(httpSecStart, httpSecStop)

	signals := system.NewSignals()
	g.Add(func() error {
		return signals.Wait(logger, g)
	}, func(error) {})

	if err := g.Run(); err != nil {
		logger.Fatal(err)
	}
}

// HttpServerConfig contains preferences for HTTP routers
type HttpServerConfig struct {
	WriteTimeOut time.Duration
	ReadTimeOut  time.Duration
	Host         string
	Port         int
}

// setupHTTP sets up HTTP server
func setupHTTP(handler http.Handler, conf *HttpServerConfig) (func() error, func(error)) {
	newSrv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		WriteTimeout: conf.WriteTimeOut,
		ReadTimeout:  conf.ReadTimeOut,
	}
	return newSrv.ListenAndServe, func(err error) { _ = newSrv.Close() }
}

// setupGRPC sets up gRPC server
func setupGRPC(baseGrpcServer *grpc.Server, basicConfig *config.BasicConfig) (func() error, func(error)) {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.GRPCPort))
	if err != nil {
		log.Fatal(err)
	}
	return func() error { return baseGrpcServer.Serve(grpcListener) }, func(err error) { _ = grpcListener.Close() }
}
