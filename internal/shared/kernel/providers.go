package kernel

import (
	"time"

	"github.com/google/uuid"
)

type Clock interface {
	Now() time.Time
}

type IDGenerator interface {
	NewID() string
}

type SystemClock struct{}

func (SystemClock) Now() time.Time {
	return time.Now().UTC()
}

type UUIDGenerator struct{}

func (UUIDGenerator) NewID() string {
	return uuid.NewString()
}
