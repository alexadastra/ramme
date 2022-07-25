package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/version"
)

// Root handler shows version
func (h *Handler) Root(w http.ResponseWriter, _ *http.Request) (int, error) {
	w.WriteHeader(http.StatusOK)
	info := make(map[string]string)
	info["service"] = config.ServiceName
	info["version"] = version.RELEASE

	resp, err := json.Marshal(info)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "failed to marshal response")
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "failed to write into response writer")
	}
	return http.StatusOK, nil
}
