package handlers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/logger"
	"github.com/alexadastra/ramme/logger/standard"
	"github.com/alexadastra/ramme/version"
)

func TestRoot(t *testing.T) {
	h := New(standard.New(&logger.Config{}), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Root)(w, r)
	})

	testHandler(t, handler, http.StatusOK,
		fmt.Sprintf("{\"service\":\"%s\",\"version\":\"%s\"}", config.ServiceName, version.RELEASE))
}
