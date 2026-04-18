package resilience

import (
	"sync"
	"time"
)

type CircuitBreaker struct {
	mu sync.Mutex

	failureThreshold int
	openDuration     time.Duration

	failureCount int
	openedAt     time.Time
}

func NewCircuitBreaker(failureThreshold int, openDuration time.Duration) *CircuitBreaker {
	if failureThreshold <= 0 {
		failureThreshold = 5
	}
	if openDuration <= 0 {
		openDuration = 30 * time.Second
	}
	return &CircuitBreaker{
		failureThreshold: failureThreshold,
		openDuration:     openDuration,
	}
}

func (c *CircuitBreaker) Allow(now time.Time) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.openedAt.IsZero() {
		return true
	}
	if now.Sub(c.openedAt) >= c.openDuration {
		c.openedAt = time.Time{}
		c.failureCount = 0
		return true
	}
	return false
}

func (c *CircuitBreaker) Success() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.failureCount = 0
	c.openedAt = time.Time{}
}

func (c *CircuitBreaker) Fail(now time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.failureCount++
	if c.failureCount >= c.failureThreshold {
		c.openedAt = now
	}
}
