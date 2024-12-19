package storage_test

import (
	"testing"

	storage "my-go-webserver/internal/storage"
)

func TestGetMessage(t *testing.T) {
	store := storage.NewStorage()

	message := store.GetMessage()
	expected := "Hello from the Go web server!"

	if message != expected {
		t.Errorf("expected %q, got %q", expected, message)
	}
}
