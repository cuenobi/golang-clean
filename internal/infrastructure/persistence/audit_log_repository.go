package persistence

import (
	"context"
	"strings"

	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	"gorm.io/gorm"
)

var _ out.AuditLogRepository = (*AuditLogRepository)(nil)

type AuditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

func (r *AuditLogRepository) List(ctx context.Context, filter out.AuditLogListFilter) ([]*entity.AuditLog, int64, error) {
	var (
		models []AuditLogModel
		total  int64
	)

	query := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Model(&AuditLogModel{})
	query = query.Where("occurred_at >= ? AND occurred_at <= ?", filter.DateFrom, filter.DateTo)

	switch filter.Scope {
	case out.AuditLogScopeSystem:
		query = query.Where("organization_id IS NULL")
	case out.AuditLogScopeOrganization:
		query = query.Where("organization_id IS NOT NULL")
	}

	if filter.ModulePrefix != "" {
		query = query.Where("module LIKE ?", filter.ModulePrefix+"%")
	}
	if len(filter.OrganizationIDs) > 0 {
		query = query.Where("organization_id IN ?", filter.OrganizationIDs)
	}
	if len(filter.Modules) > 0 {
		query = query.Where("module IN ?", filter.Modules)
	}
	if len(filter.Actions) > 0 {
		query = query.Where("action IN ?", filter.Actions)
	}
	if len(filter.Usernames) > 0 {
		query = query.Where("username IN ?", filter.Usernames)
	}
	if len(filter.EntityIDs) > 0 {
		query = query.Where("entity_id IN ?", filter.EntityIDs)
	}
	if len(filter.EntityTypes) > 0 {
		query = query.Where("entity_type IN ?", filter.EntityTypes)
	}
	if search := strings.TrimSpace(filter.Search); search != "" {
		term := "%" + search + "%"
		query = query.Where(
			"(username ILIKE ? OR module ILIKE ? OR action ILIKE ? OR entity_type ILIKE ? OR COALESCE(entity_name, '') ILIKE ?)",
			term,
			term,
			term,
			term,
			term,
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := sanitizeAuditLogSortBy(filter.SortBy)
	sortOrder := sanitizeAuditLogSortOrder(filter.SortOrder)
	offset := (filter.Page - 1) * filter.PageSize
	if offset < 0 {
		offset = 0
	}

	if err := query.Order(sortBy + " " + sortOrder).Offset(offset).Limit(filter.PageSize).Find(&models).Error; err != nil {
		return nil, 0, err
	}

	result := make([]*entity.AuditLog, 0, len(models))
	for _, model := range models {
		result = append(result, toAuditLogEntity(model))
	}

	return result, total, nil
}

func toAuditLogEntity(model AuditLogModel) *entity.AuditLog {
	entityName := ""
	if model.EntityName != nil {
		entityName = *model.EntityName
	}
	ipAddress := ""
	if model.IPAddress != nil {
		ipAddress = *model.IPAddress
	}
	userAgent := ""
	if model.UserAgent != nil {
		userAgent = *model.UserAgent
	}
	organizationName := ""
	if model.OrganizationName != nil {
		organizationName = *model.OrganizationName
	}

	return &entity.AuditLog{
		ID:               model.ID,
		EventID:          model.EventID,
		EntityType:       model.EntityType,
		EntityID:         model.EntityID,
		EntityName:       entityName,
		Action:           model.Action,
		Username:         model.Username,
		Module:           model.Module,
		IPAddress:        ipAddress,
		UserAgent:        userAgent,
		DiffValue:        model.DiffValue,
		OrganizationID:   model.OrganizationID,
		OrganizationName: organizationName,
		OccurredAt:       model.OccurredAt.UTC(),
		CreatedAt:        model.CreatedAt.UTC(),
	}
}

func sanitizeAuditLogSortBy(sortBy string) string {
	switch strings.TrimSpace(strings.ToLower(sortBy)) {
	case "id":
		return "id"
	case "created_at":
		return "created_at"
	case "occurred_at":
		return "occurred_at"
	case "module":
		return "module"
	case "action":
		return "action"
	case "username":
		return "username"
	case "entity_type":
		return "entity_type"
	default:
		return "occurred_at"
	}
}

func sanitizeAuditLogSortOrder(sortOrder string) string {
	if strings.EqualFold(strings.TrimSpace(sortOrder), "asc") {
		return "ASC"
	}
	return "DESC"
}
