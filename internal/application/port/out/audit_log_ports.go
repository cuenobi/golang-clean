package out

import (
	"context"
	"time"

	"github.com/cuenobi/golang-clean/internal/domain/entity"
)

type AuditLogScope string

const (
	AuditLogScopeSystem       AuditLogScope = "system"
	AuditLogScopeOrganization AuditLogScope = "organization"
)

type AuditLogListFilter struct {
	Scope           AuditLogScope
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
	ModulePrefix    string
}

type AuditLogRepository interface {
	List(ctx context.Context, filter AuditLogListFilter) ([]*entity.AuditLog, int64, error)
}
