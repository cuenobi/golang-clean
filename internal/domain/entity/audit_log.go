package entity

import "time"

// AuditLog stores immutable audit event data for API queries.
type AuditLog struct {
	ID               uint64
	EventID          string
	EntityType       string
	EntityID         int64
	EntityName       string
	Action           string
	Username         string
	Module           string
	IPAddress        string
	UserAgent        string
	DiffValue        []byte
	OrganizationID   *int64
	OrganizationName string
	OccurredAt       time.Time
	CreatedAt        time.Time
}
