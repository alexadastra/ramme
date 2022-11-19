package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexadastra/ramme/config_new"
	"github.com/alexadastra/ramme/handlers"
)

func TestSetup(t *testing.T) {
	conf, st, end, err := config_new.NewConfig("./config_new/tests/config.yaml")
	if err != nil {
		panic(err)
	}
	go st()
	defer end()
	router, logger, err := Setup(conf)
	if err != nil {
		t.Errorf("Fail, got '%s', want '%v'", err, nil)
	}
	if router == nil {
		t.Error("Expected new router, got nil")
	}
	if logger == nil {
		t.Error("Expected new logger, got nil")
	}

	h := handlers.New(logger, conf)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(notFound)(w, r)
	})

	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Error(err)
	}

	trw := httptest.NewRecorder()
	handler.ServeHTTP(trw, req)

	if trw.Code != http.StatusNotFound {
		t.Error("Expected status:", http.StatusNotFound, "got", trw.Code)
	}
}
