package http

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	useCase in.AuditLogUseCase
}

func NewHandler(useCase in.AuditLogUseCase) *Handler {
	return &Handler{useCase: useCase}
}

// GetSystemAuditLogs godoc
// @Summary Get system audit logs
// @Description Get system-level audit logs with filters and pagination.
// @Tags Audit Logs
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: audit_logs:read"
// @Param date_from query string true "Start date in RFC3339"
// @Param date_to query string true "End date in RFC3339"
// @Param modules query []string false "Module filters"
// @Param actions query []string false "Action filters"
// @Param usernames query []string false "Username filters"
// @Param entity_ids query []int false "Entity ID filters"
// @Param entity_types query []string false "Entity type filters"
// @Param search query string false "Free text search"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} ListAuditLogsResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/audit-logs/system [get]
func (h *Handler) GetSystemAuditLogs(c *fiber.Ctx) error {
	req, err := toListAuditLogsRequest(c)
	if err != nil {
		return err
	}

	result, err := h.useCase.GetSystemAuditLogs(c.UserContext(), req)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// GetOrganizationAuditLogs godoc
// @Summary Get organization audit logs
// @Description Get organization-level audit logs with filters and pagination.
// @Tags Audit Logs
// @Security ApiKeyAuth
// @Produce json
// @Param X-Permissions header string true "Required permission: audit_logs:read"
// @Param date_from query string true "Start date in RFC3339"
// @Param date_to query string true "End date in RFC3339"
// @Param organization_ids query []int false "Organization ID filters"
// @Param modules query []string false "Module filters"
// @Param actions query []string false "Action filters"
// @Param usernames query []string false "Username filters"
// @Param entity_ids query []int false "Entity ID filters"
// @Param entity_types query []string false "Entity type filters"
// @Param search query string false "Free text search"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} ListAuditLogsResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 401 {object} ErrorResponseDoc
// @Failure 403 {object} ErrorResponseDoc
// @Failure 429 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /api/v1/audit-logs/organization [get]
func (h *Handler) GetOrganizationAuditLogs(c *fiber.Ctx) error {
	req, err := toListAuditLogsRequest(c)
	if err != nil {
		return err
	}

	result, err := h.useCase.GetOrganizationAuditLogs(c.UserContext(), req)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
