package di

import (
	"github.com/IBM/sarama"
	messaginginfra "github.com/cuenobi/golang-clean/internal/infrastructure/messaging"
	"github.com/cuenobi/golang-clean/internal/shared/kafkax"
)

func (c *Container) wireMessaging() error {
	kafkaConfig := kafkax.NewDefaultSaramaConfig(c.Cfg)
	producer, err := sarama.NewSyncProducer(c.Cfg.KafkaBrokers, kafkaConfig)
	if err != nil {
		return err
	}

	c.Producer = producer
	c.EventPublisher = messaginginfra.NewPublisher(producer, "order.created.v1")
	return nil
}
