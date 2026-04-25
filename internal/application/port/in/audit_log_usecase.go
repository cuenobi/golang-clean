package in

import (
	"context"

	auditlogdto "github.com/cuenobi/golang-clean/internal/application/dto/auditlog"
)

type AuditLogUseCase interface {
	GetSystemAuditLogs(ctx context.Context, req auditlogdto.ListAuditLogsRequest) (auditlogdto.ListAuditLogsResponse, error)
	GetOrganizationAuditLogs(ctx context.Context, req auditlogdto.ListAuditLogsRequest) (auditlogdto.ListAuditLogsResponse, error)
}
