// Package service defines application inits and start ups
package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/alexadastra/ramme/config_new"
	"github.com/alexadastra/ramme/system"
	"google.golang.org/grpc"
)

// Run inits and starts up ramme app components
func Run(ctx context.Context, g *system.GroupOperator, baseGrpcServer *grpc.Server, mux http.Handler, conf *config_new.Config) {
	host := config_new.ToString(conf.Get(config_new.Host))
	grpcPort := config_new.ToInt(conf.Get(config_new.GRPCPort))
	httpPort := config_new.ToInt(conf.Get(config_new.HTTPPort))
	httpPortSec := config_new.ToInt(conf.Get(config_new.HTTPSecondaryPort))

	// Configure service and get router
	router, logger, err := Setup(conf)
	if err != nil {
		log.Fatal(err)
	}

	grpcStart, grpcStop := setupGRPC(baseGrpcServer, host, grpcPort)
	g.Add(grpcStart, grpcStop)
	logger.Warnf("Serving grpc address %s", fmt.Sprintf("%s:%d", host, grpcPort))

	httpStart, httpStop := setupHTTP(mux, &HTTPServerConfig{
		WriteTimeOut: 15 * time.Second,
		ReadTimeOut:  15 * time.Second,
		Host:         host,
		Port:         httpPort,
	})
	g.Add(httpStart, httpStop)
	logger.Warnf("Serving http address %s", fmt.Sprintf("%s:%d", host, httpPort))

	httpSecStart, httpSecStop := setupHTTP(router, &HTTPServerConfig{
		WriteTimeOut: 15 * time.Second,
		ReadTimeOut:  15 * time.Second,
		Host:         host,
		Port:         httpPortSec,
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

// HTTPServerConfig contains preferences for HTTP routers
type HTTPServerConfig struct {
	WriteTimeOut time.Duration
	ReadTimeOut  time.Duration
	Host         string
	Port         int
}

// setupHTTP sets up HTTP server
func setupHTTP(handler http.Handler, conf *HTTPServerConfig) (func() error, func(error)) {
	newSrv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		WriteTimeout: conf.WriteTimeOut,
		ReadTimeout:  conf.ReadTimeOut,
	}
	return newSrv.ListenAndServe, func(err error) { _ = newSrv.Close() }
}

// setupGRPC sets up gRPC server
func setupGRPC(baseGrpcServer *grpc.Server, host string, port int) (func() error, func(error)) {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
	return func() error { return baseGrpcServer.Serve(grpcListener) }, func(err error) { _ = grpcListener.Close() }
}
