package usecase

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
)

const (
	systemModulePrefix       = "SYSTEM_"
	organizationModulePrefix = "ORG_"
)

var _ in.AuditLogUseCase = (*AuditLogUseCase)(nil)

type AuditLogUseCase struct {
	repo out.AuditLogRepository
}

func NewAuditLogUseCase(repo out.AuditLogRepository) *AuditLogUseCase {
	return &AuditLogUseCase{repo: repo}
}
