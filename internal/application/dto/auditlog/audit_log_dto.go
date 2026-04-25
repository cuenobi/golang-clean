package auditlogdto

import (
	"strings"
	"time"
)

const (
	defaultAuditLogPage     = 1
	defaultAuditLogPageSize = 20
	maxAuditLogPageSize     = 100
)

type ListAuditLogsRequest struct {
	OrganizationIDs []int64
	DateFrom        time.Time
	DateTo          time.Time
	Modules         []string
	Actions         []string
	Usernames       []string
	EntityIDs       []int64
	EntityTypes     []string
	Search          string
	Page            int
	PageSize        int
	SortBy          string
	SortOrder       string
}

func (r *ListAuditLogsRequest) ApplyDefaults() {
	if r.Page <= 0 {
		r.Page = defaultAuditLogPage
	}
	if r.PageSize <= 0 {
		r.PageSize = defaultAuditLogPageSize
	}
	if r.PageSize > maxAuditLogPageSize {
		r.PageSize = maxAuditLogPageSize
	}
	if strings.TrimSpace(r.SortBy) == "" {
		r.SortBy = "occurred_at"
	}
	if strings.TrimSpace(r.SortOrder) == "" {
		r.SortOrder = "DESC"
	}
}

type AuditLogResponse struct {
	ID               uint64     `json:"id"`
	EventID          string     `json:"event_id"`
	EntityType       string     `json:"entity_type"`
	EntityID         int64      `json:"entity_id"`
	EntityName       string     `json:"entity_name,omitempty"`
	Action           string     `json:"action"`
	Username         string     `json:"username"`
	Module           string     `json:"module"`
	IPAddress        string     `json:"ip_address,omitempty"`
	UserAgent        string     `json:"user_agent,omitempty"`
	DiffValue        any        `json:"diff_value,omitempty"`
	OrganizationID   *int64     `json:"organization_id,omitempty"`
	OrganizationName string     `json:"organization_name,omitempty"`
	OccurredAt       time.Time  `json:"occurred_at"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
}

type ListAuditLogsResponse struct {
	Data       []AuditLogResponse `json:"data"`
	Total      int64              `json:"total"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
	TotalPages int                `json:"total_pages"`
}
