package service

import (
	"fmt"
	"net/http"

	"github.com/alexadastra/ramme/handlers"

	"github.com/pkg/errors"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/logger"
	stdlog "github.com/alexadastra/ramme/logger/standard"
	"github.com/alexadastra/ramme/version"
	"github.com/gorilla/mux"
)

// Setup configures the service
func Setup(cfg *config.BasicConfig) (*mux.Router, logger.Logger, error) {
	// Setup logger
	l := stdlog.New(&logger.Config{
		Level: cfg.LogLevel,
		Time:  true,
		UTC:   true,
	})

	l.Info("Version:", version.RELEASE)
	l.Warnf("%s log level is used", cfg.LogLevel.String())
	l.Infof("Service %s listens secondary requests on %s:%d", config.ServiceName, cfg.Host, cfg.HTTPSecondaryPort)

	// Define handlers
	h := handlers.New(l, cfg)

	// Register new router
	r := mux.NewRouter()

	// Response for undefined methods
	r.NotFoundHandler = http.HandlerFunc(h.Base(notFound))

	r.HandleFunc("/", h.Base(h.Root)).Methods("GET")
	r.HandleFunc("/healthz", h.Base(h.Health)).Methods("GET")
	r.HandleFunc("/readyz", h.Base(h.Ready)).Methods("GET")
	r.HandleFunc("/info", h.Base(h.Info)).Methods("GET")

	return r, l, nil
}

// Response for undefined methods
func notFound(w http.ResponseWriter, r *http.Request) (int, error) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(fmt.Sprintf("Method not found for %s", r.URL.Path)))
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "failed to write into response writer")
	}
	return http.StatusNotFound, nil
}
