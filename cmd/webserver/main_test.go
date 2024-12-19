package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"my-go-webserver/internal/handler"
	"my-go-webserver/internal/service"
	"my-go-webserver/internal/storage"
)

func TestMainHomeHandler(t *testing.T) {
	storageLayer := storage.NewStorage()
	serviceLayer := service.NewService(storageLayer)
	handlerLayer := handler.NewHandler(serviceLayer)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handlerLayer.Home(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expected := "Hello from the Go web server!"
	if w.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, w.Body.String())
	}
}
