package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go-kafka-project/shared"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "events",
	})

	defer writer.Close()

	for i := 1; i <= 10; i++ {
		event := shared.Event{
			ID:    string(rune(i)),
			Type:  "sensor",
			Value: "temperature",
			Time:  time.Now().String(),
		}

		data, _ := json.Marshal(event)

		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(event.ID),
				Value: data,
			},
		)

		if err != nil {
			log.Fatal("failed to write message:", err)
		}

		log.Println("Sent:", string(data))
		time.Sleep(1 * time.Second)
	}
}