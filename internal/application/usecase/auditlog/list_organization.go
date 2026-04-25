package usecase

import (
	"context"

	auditlogdto "github.com/cuenobi/golang-clean/internal/application/dto/auditlog"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
)

func (u *AuditLogUseCase) GetOrganizationAuditLogs(ctx context.Context, req auditlogdto.ListAuditLogsRequest) (auditlogdto.ListAuditLogsResponse, error) {
	req.ApplyDefaults()
	req.Modules = filterModulesByPrefix(req.Modules, organizationModulePrefix)

	logs, total, err := u.repo.List(ctx, toFilter(req, out.AuditLogScopeOrganization, organizationModulePrefix))
	if err != nil {
		return auditlogdto.ListAuditLogsResponse{}, err
	}

	return toListResponse(logs, total, req), nil
}
