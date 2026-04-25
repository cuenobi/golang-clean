package persistence

import "time"

type AuditLogModel struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	EventID          string    `gorm:"column:event_id;not null;index:idx_audit_log_event_id"`
	EntityType       string    `gorm:"column:entity_type;not null;index:idx_audit_log_entity"`
	EntityID         int64     `gorm:"column:entity_id;not null;index:idx_audit_log_entity"`
	EntityName       *string   `gorm:"column:entity_name"`
	Action           string    `gorm:"column:action;not null;index:idx_audit_log_action"`
	Username         string    `gorm:"column:username;not null;index:idx_audit_log_username"`
	Module           string    `gorm:"column:module;not null;index:idx_audit_log_module"`
	IPAddress        *string   `gorm:"column:ip_address"`
	UserAgent        *string   `gorm:"column:user_agent"`
	DiffValue        []byte    `gorm:"column:diff_value;type:jsonb"`
	OrganizationID   *int64    `gorm:"column:organization_id;index:idx_audit_log_organization_id"`
	OrganizationName *string   `gorm:"column:organization_name"`
	OccurredAt       time.Time `gorm:"column:occurred_at;not null;index:idx_audit_log_occurred_at"`
	CreatedAt        time.Time `gorm:"column:created_at;not null"`
}

func (AuditLogModel) TableName() string {
	return "audit_log"
}
