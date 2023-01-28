package handlers

import (
	"net/http"
	"testing"

	"github.com/alexadastra/ramme/logger"
	"github.com/alexadastra/ramme/logger/standard"
)

func TestReady(t *testing.T) {
	h := New(standard.New(&logger.Config{}), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Ready)(w, r)
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
