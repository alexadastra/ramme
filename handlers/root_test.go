package handlers

import (
	"fmt"
	"net/http"
	"testing"

	"git.miem.hse.ru/786/ramme/config"
	"git.miem.hse.ru/786/ramme/logger"
	"git.miem.hse.ru/786/ramme/logger/standard"
	"git.miem.hse.ru/786/ramme/version"
)

func TestRoot(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Root)(w, r)
	})

	testHandler(t, handler, http.StatusOK,
		fmt.Sprintf("{\"service\":\"%s\",\"version\":\"%s\"}", config.SERVICENAME, version.RELEASE))
}
