package service_test

import (
	"testing"

	"my-go-webserver/internal/service"
	"my-go-webserver/internal/storage"
)

func TestGetWelcomeMessage(t *testing.T) {
	mockStorage := storage.NewStorage()
	serviceInstance := service.NewService(mockStorage)

	message := serviceInstance.GetWelcomeMessage()
	expected := "Hello from the Go web server!"

	if message != expected {
		t.Errorf("expected %q, got %q", expected, message)
	}
}
