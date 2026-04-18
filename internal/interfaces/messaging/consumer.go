package messaging

import (
	"context"
	"time"

	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/cuenobi/golang-clean/internal/shared/logger"
	"github.com/cuenobi/golang-clean/internal/shared/resilience"
)

type Consumer struct {
	logger    logger.Logger
	outbox    out.OutboxStore
	publisher out.EventPublisher
	cfg       config.Config
	clock     out.Clock
	breaker   *resilience.CircuitBreaker
}

func NewConsumer(
	log logger.Logger,
	outbox out.OutboxStore,
	publisher out.EventPublisher,
	cfg config.Config,
	clock out.Clock,
) *Consumer {
	openDuration := time.Duration(cfg.CircuitBreakerOpenMS) * time.Millisecond
	return &Consumer{
		logger:    log,
		outbox:    outbox,
		publisher: publisher,
		cfg:       cfg,
		clock:     clock,
		breaker:   resilience.NewCircuitBreaker(cfg.CircuitBreakerFailures, openDuration),
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	c.logger.Info("consumer_loop_started", map[string]any{
		"component": "order_consumer",
	})

	pollEvery := time.Duration(c.cfg.OutboxPollIntervalMS) * time.Millisecond
	if pollEvery <= 0 {
		pollEvery = time.Second
	}

	ticker := time.NewTicker(pollEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := c.dispatchBatch(ctx); err != nil {
				c.logger.Error("outbox_dispatch_batch_failed", err, nil)
			}
		}
	}
}

func (c *Consumer) dispatchBatch(ctx context.Context) error {
	now := c.clock.Now()
	if !c.breaker.Allow(now) {
		c.logger.Warn("outbox_dispatch_paused_by_circuit_breaker", map[string]any{
			"component": "order_consumer",
		})
		return nil
	}

	processingTimeout := time.Duration(c.cfg.OutboxProcessingTimeoutMS) * time.Millisecond
	messages, err := c.outbox.ClaimPending(ctx, now, c.cfg.OutboxBatchSize, processingTimeout)
	if err != nil {
		return err
	}
	if len(messages) == 0 {
		return nil
	}

	for _, msg := range messages {
		if err := c.publishOne(ctx, msg); err != nil {
			c.logger.Error("outbox_publish_failed", err, map[string]any{
				"message_id":  msg.ID,
				"event_type":  msg.EventType,
				"retry_count": msg.RetryCount,
			})
		}
	}
	return nil
}

func (c *Consumer) publishOne(ctx context.Context, msg out.OutboxMessage) error {
	now := c.clock.Now()

	timeout := time.Duration(c.cfg.KafkaPublishTimeoutMS) * time.Millisecond
	if timeout <= 0 {
		timeout = 3 * time.Second
	}
	pubCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := c.publisher.PublishOrderCreated(pubCtx, msg.Payload)
	if err == nil {
		c.breaker.Success()
		return c.outbox.MarkPublished(ctx, msg.ID, now)
	}

	c.breaker.Fail(now)
	retryCount := msg.RetryCount + 1
	dead := retryCount >= c.cfg.OutboxMaxRetries && c.cfg.OutboxMaxRetries > 0
	nextRetryAt := now.Add(c.retryBackoff(retryCount))
	if dead {
		nextRetryAt = now
	}

	if markErr := c.outbox.MarkRetry(ctx, msg.ID, retryCount, nextRetryAt, err.Error(), dead); markErr != nil {
		return markErr
	}
	return err
}

func (c *Consumer) retryBackoff(retryCount int) time.Duration {
	base := time.Duration(c.cfg.OutboxRetryBackoffMS) * time.Millisecond
	if base <= 0 {
		base = 500 * time.Millisecond
	}
	if retryCount <= 1 {
		return base
	}

	shift := retryCount - 1
	if shift > 6 {
		shift = 6
	}
	return base * time.Duration(1<<shift)
}
