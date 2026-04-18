package messaging

import (
	"context"

	"github.com/cuenobi/golang-clean/internal/shared/logger"
)

type Consumer struct {
	logger logger.Logger
}

func NewConsumer(log logger.Logger) *Consumer {
	return &Consumer{logger: log}
}

func (c *Consumer) Run(ctx context.Context) error {
	c.logger.Info("consumer_loop_started", map[string]any{
		"component": "order_consumer",
	})
	<-ctx.Done()
	return ctx.Err()
}
