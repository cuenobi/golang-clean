package kafkax

import "github.com/IBM/sarama"

func NewDefaultSaramaConfig(clientID string) *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.ClientID = clientID
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Consumer.Return.Errors = true
	cfg.Version = sarama.V3_6_0_0
	return cfg
}
