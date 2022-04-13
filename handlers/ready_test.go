package handlers

import (
	"net/http"
	"testing"

	"git.miem.hse.ru/786/ramme/config"
	"git.miem.hse.ru/786/ramme/logger"
	"git.miem.hse.ru/786/ramme/logger/standard"
)

func TestReady(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Ready)(w, r)
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
