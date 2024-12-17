package config

import (
	"os"
	"log"
)

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Env variable %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
