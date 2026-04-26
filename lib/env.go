package lib

import (
	"os"
)

var (
	NatsURL = Getenv("NATS_URL", "nats://rhel10:4222")
)

func Getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
