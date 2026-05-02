package main

import (
	"context"
	"log"

	"go-kafka-microservices/internal/config"
	"go-kafka-microservices/internal/kafka"
)

func main() {
	reader := kafka.NewReader(config.KafkaBroker(), "orders", "analytics-group")
	defer reader.Close()

	log.Println("Analytics Service started...")

	for {
		msg, err := kafka.Read(context.Background(), reader)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("Analyzing: %s\n", string(msg.Value))
	}
}