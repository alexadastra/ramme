package handlers

import (
	"net/http"

	"github.com/pkg/errors"
)

// Health returns "OK" if service is alive
func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) (int, error) {
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/text")
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "failed to write into response writer")
	}
	return http.StatusOK, nil
}
