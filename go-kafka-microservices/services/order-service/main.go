package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"go-kafka-microservices/internal/config"
	"go-kafka-microservices/internal/kafka"
	"go-kafka-microservices/internal/models"
)

func main() {
	broker := config.KafkaBroker()
	log.Println("Using Kafka broker:", broker)

	writer := kafka.NewWriter(broker, "orders")
	defer writer.Close()

	log.Println("Order Service started...")

	for {
		event := models.Event{
			ID:        uuid.NewString(),
			Type:      "order_created",
			Payload:   "order_ABC",
			Timestamp: time.Now().Format(time.RFC3339),
		}

		data, err := json.Marshal(event)
		if err != nil {
			log.Println("json error:", err)
			continue
		}

		err = kafka.Write(context.Background(), writer, []byte(event.ID), data)
		if err != nil {
			log.Println("Kafka write error:", err)
		}

		log.Println("Produced:", string(data))
		time.Sleep(2 * time.Second)
	}
}