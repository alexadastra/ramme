// Package handlers defines admin HTTP app handlers
package handlers

import (
	"net/http"
	"time"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/logger"
)

// Handler defines common part for all handlers
type Handler struct {
	logger      logger.Logger
	config      *config.Config
	maintenance bool
	stats       *stats
}

// TODO: stats should also be linked with gRPC and HTTP handlers
type stats struct {
	requests        *Requests
	averageDuration time.Duration
	maxDuration     time.Duration
	totalDuration   time.Duration
	requestsCount   time.Duration
	startTime       time.Time
}

// New returns new instance of the Handler
func New(logger logger.Logger, conf *config.Config) *Handler {
	return &Handler{
		logger: logger,
		config: conf,
		stats: &stats{
			requests:  new(Requests),
			startTime: time.Now(),
		},
	}
}

// Base handler implements middleware logic
func (h *Handler) Base(handle func(w http.ResponseWriter, r *http.Request) (int,
	error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := time.Now()
		code, err := handle(w, r)
		if err != nil {
			h.logger.Errorf("failed to process request: %s", err)
		}
		h.countDuration(timer)
		h.collectCodes(code)
	}
}

func (h *Handler) countDuration(timer time.Time) {
	if !timer.IsZero() {
		h.stats.requestsCount++
		took := time.Now()
		duration := took.Sub(timer)
		h.stats.totalDuration += duration
		if duration > h.stats.maxDuration {
			h.stats.maxDuration = duration
		}
		h.stats.averageDuration = h.stats.totalDuration / h.stats.requestsCount
		h.stats.requests.Duration.Max = h.stats.maxDuration.String()
		h.stats.requests.Duration.Average = h.stats.averageDuration.String()
	}
}

func (h *Handler) collectCodes(code int) {
	if code >= 500 {
		h.stats.requests.Codes.C5xx++
	} else {
		if code >= 400 {
			h.stats.requests.Codes.C4xx++
		} else if code >= 200 && code < 300 {
			h.stats.requests.Codes.C2xx++
		}
	}
}
