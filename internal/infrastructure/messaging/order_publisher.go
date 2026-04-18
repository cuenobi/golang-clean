package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
)

var _ out.EventPublisher = (*Publisher)(nil)

type Publisher struct {
	producer sarama.SyncProducer
	topic    string
}

func NewPublisher(producer sarama.SyncProducer, topic string) *Publisher {
	return &Publisher{producer: producer, topic: topic}
}

func (p *Publisher) PublishOrderCreated(ctx context.Context, eventPayload any) error {
	payload, err := json.Marshal(eventPayload)
	if err != nil {
		return err
	}
	_, _, err = p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(payload),
	})
	if err != nil {
		return fmt.Errorf("publish to kafka topic %s: %w", p.topic, err)
	}
	return nil
}
