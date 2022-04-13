package handlers

import (
	"net/http"

	"github.com/pkg/errors"
)

// Ready returns "OK" if service is ready to serve traffic
func (h *Handler) Ready(w http.ResponseWriter, _ *http.Request) (int, error) {
	// TODO: possible use cases:
	// load data from a database, a message broker, any external services, etc

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/text")
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "failed to write into response writer")
	}
	return http.StatusOK, nil
}
