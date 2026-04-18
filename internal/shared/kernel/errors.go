package kernel

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidState = errors.New("invalid state")
	ErrConflict     = errors.New("conflict")
)
