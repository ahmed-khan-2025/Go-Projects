package kafka

import (
	"context"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

func NewWriter(broker, topic string) *kafkago.Writer {
	return &kafkago.Writer{
		Addr:     kafkago.TCP(broker),
		Topic:    topic,
		Balancer: &kafkago.LeastBytes{},
		Async:    false,
	}
}

func NewReader(broker, topic, groupID string) *kafkago.Reader {
	return kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafkago.FirstOffset,
	})
}

func Write(ctx context.Context, w *kafkago.Writer, key, value []byte) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return w.WriteMessages(ctx, kafkago.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	})
}

func Read(ctx context.Context, r *kafkago.Reader) (kafkago.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return r.ReadMessage(ctx)
}