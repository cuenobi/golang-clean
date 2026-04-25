package usecase

import (
	"encoding/json"
	"strings"

	auditlogdto "github.com/cuenobi/golang-clean/internal/application/dto/auditlog"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
)

func toFilter(req auditlogdto.ListAuditLogsRequest, scope out.AuditLogScope, modulePrefix string) out.AuditLogListFilter {
	return out.AuditLogListFilter{
		Scope:           scope,
		OrganizationIDs: req.OrganizationIDs,
		DateFrom:        req.DateFrom,
		DateTo:          req.DateTo,
		Modules:         req.Modules,
		Actions:         req.Actions,
		Usernames:       req.Usernames,
		EntityIDs:       req.EntityIDs,
		EntityTypes:     req.EntityTypes,
		Search:          strings.TrimSpace(req.Search),
		Page:            req.Page,
		PageSize:        req.PageSize,
		SortBy:          req.SortBy,
		SortOrder:       req.SortOrder,
		ModulePrefix:    modulePrefix,
	}
}

func toListResponse(logs []*entity.AuditLog, total int64, req auditlogdto.ListAuditLogsRequest) auditlogdto.ListAuditLogsResponse {
	items := make([]auditlogdto.AuditLogResponse, 0, len(logs))
	for _, log := range logs {
		items = append(items, toResponse(log))
	}

	totalPages := 0
	if req.PageSize > 0 {
		totalPages = int((total + int64(req.PageSize) - 1) / int64(req.PageSize))
	}

	return auditlogdto.ListAuditLogsResponse{
		Data:       items,
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
	}
}

func toResponse(log *entity.AuditLog) auditlogdto.AuditLogResponse {
	var diff any
	if len(log.DiffValue) > 0 {
		_ = json.Unmarshal(log.DiffValue, &diff)
	}

	createdAt := log.CreatedAt

	return auditlogdto.AuditLogResponse{
		ID:               log.ID,
		EventID:          log.EventID,
		EntityType:       log.EntityType,
		EntityID:         log.EntityID,
		EntityName:       log.EntityName,
		Action:           log.Action,
		Username:         log.Username,
		Module:           log.Module,
		IPAddress:        log.IPAddress,
		UserAgent:        log.UserAgent,
		DiffValue:        diff,
		OrganizationID:   log.OrganizationID,
		OrganizationName: log.OrganizationName,
		OccurredAt:       log.OccurredAt,
		CreatedAt:        &createdAt,
	}
}

func filterModulesByPrefix(modules []string, prefix string) []string {
	if len(modules) == 0 {
		return modules
	}

	result := make([]string, 0, len(modules))
	for _, module := range modules {
		clean := strings.TrimSpace(module)
		if clean == "" {
			continue
		}
		if strings.HasPrefix(clean, prefix) {
			result = append(result, clean)
		}
	}

	return result
}
