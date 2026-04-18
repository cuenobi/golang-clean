package kafkax

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/cuenobi/golang-clean/internal/shared/config"
)

func NewDefaultSaramaConfig(appCfg config.Config) *sarama.Config {
	saramaCfg := sarama.NewConfig()
	saramaCfg.ClientID = appCfg.AppName
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Retry.Max = 3
	saramaCfg.Producer.Retry.Backoff = 200 * time.Millisecond
	saramaCfg.Producer.Timeout = 5 * time.Second

	saramaCfg.Net.DialTimeout = 5 * time.Second
	saramaCfg.Net.ReadTimeout = 5 * time.Second
	saramaCfg.Net.WriteTimeout = 5 * time.Second

	saramaCfg.Consumer.Return.Errors = true
	saramaCfg.Version = sarama.V3_6_0_0
	return saramaCfg
}
