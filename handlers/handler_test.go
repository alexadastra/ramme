package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexadastra/ramme/logger"
	"github.com/alexadastra/ramme/logger/standard"
)

func TestCollectCodes(t *testing.T) {
	h := New(standard.New(&logger.Config{}), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(func(w http.ResponseWriter, r *http.Request) (int, error) {
			w.WriteHeader(http.StatusBadGateway)
			_, _ = w.Write([]byte(http.StatusText(http.StatusBadGateway)))
			return http.StatusBadGateway, nil
		})(w, r)
	})
	testHandler(t, handler, http.StatusBadGateway, http.StatusText(http.StatusBadGateway))

	handler = func(w http.ResponseWriter, r *http.Request) {
		h.Base(func(w http.ResponseWriter, r *http.Request) (int, error) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return http.StatusNotFound, nil
		})(w, r)
	}
	testHandler(t, handler, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func testHandler(t *testing.T, handler http.HandlerFunc, code int, body string) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	trw := httptest.NewRecorder()
	handler.ServeHTTP(trw, req)

	if trw.Code != code {
		t.Error("Expected status code:", code, "got", trw.Code)
	}
	if trw.Body.String() != body {
		t.Error("Expected body", body, "got", trw.Body.String())
	}
}
