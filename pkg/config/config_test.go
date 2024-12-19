package config_test

import (
	"os"
	"testing"

	"my-go-webserver/pkg/config"
)

func TestGetEnv(t *testing.T) {
	key := "TEST_ENV_VAR"
	defaultValue := "default"

	os.Unsetenv(key)
	result := config.GetEnv(key, defaultValue)
	if result != defaultValue {
		t.Errorf("expected %q, got %q", defaultValue, result)
	}

	expectedValue := "value"
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key)

	result = config.GetEnv(key, defaultValue)
	if result != expectedValue {
		t.Errorf("expected %q, got %q", expectedValue, result)
	}
}
