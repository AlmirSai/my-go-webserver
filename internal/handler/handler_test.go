package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"my-go-webserver/internal/handler"
	"my-go-webserver/internal/service"
	"my-go-webserver/internal/storage"
)

func TestHome(t *testing.T) {
	mockStorage := storage.NewStorage()
	mockService := service.NewService(mockStorage)
	h := handler.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.Home(w, req)

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
