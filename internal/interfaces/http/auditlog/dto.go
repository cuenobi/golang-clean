package http

import (
	auditlogdto "github.com/cuenobi/golang-clean/internal/application/dto/auditlog"
	sharedhttpx "github.com/cuenobi/golang-clean/internal/shared/httpx"
)

type ErrorResponseDoc = sharedhttpx.ErrorResponse
type ListAuditLogsResponseDoc = auditlogdto.ListAuditLogsResponse

type listAuditLogsQuery struct {
	DateFrom  string `json:"date_from" validate:"required"`
	DateTo    string `json:"date_to" validate:"required"`
	Page      int    `json:"page" validate:"omitempty,min=1"`
	PageSize  int    `json:"page_size" validate:"omitempty,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc ASC desc DESC"`
}
