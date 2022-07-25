package handlers

import (
	"net/http"
	"testing"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/logger"
	"github.com/alexadastra/ramme/logger/standard"
)

func TestHealth(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.BasicConfig))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Health)(w, r)
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
