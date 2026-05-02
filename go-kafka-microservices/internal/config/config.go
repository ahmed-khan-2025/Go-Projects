package config

import "os"

func KafkaBroker() string {
	if v := os.Getenv("KAFKA_BROKER"); v != "" {
		return v
	}
	return "localhost:29092"
}