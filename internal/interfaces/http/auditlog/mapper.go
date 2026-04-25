package http

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	auditlogdto "github.com/cuenobi/golang-clean/internal/application/dto/auditlog"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedvalidator "github.com/cuenobi/golang-clean/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
)

func toListAuditLogsRequest(c *fiber.Ctx) (auditlogdto.ListAuditLogsRequest, error) {
	values, err := url.ParseQuery(string(c.Context().URI().QueryString()))
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, kernel.NewBadRequestError("invalid query string")
	}

	page, err := parseOptionalInt(values.Get("page"), "page")
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}
	pageSize, err := parseOptionalInt(values.Get("page_size"), "page_size")
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}

	query := listAuditLogsQuery{
		DateFrom:  strings.TrimSpace(values.Get("date_from")),
		DateTo:    strings.TrimSpace(values.Get("date_to")),
		Page:      page,
		PageSize:  pageSize,
		SortBy:    strings.TrimSpace(values.Get("sort_by")),
		SortOrder: strings.TrimSpace(values.Get("sort_order")),
	}
	if err := sharedvalidator.ValidateStruct(query); err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}

	dateFrom, err := parseRFC3339(query.DateFrom, "date_from")
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}
	dateTo, err := parseRFC3339(query.DateTo, "date_to")
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}
	if dateTo.Before(dateFrom) {
		return auditlogdto.ListAuditLogsRequest{}, newQueryValidationError("date_to", "gtefield", "date_to must be greater than or equal to date_from")
	}

	organizationIDs, err := parseInt64List(values["organization_ids"], "organization_ids")
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}
	entityIDs, err := parseInt64List(values["entity_ids"], "entity_ids")
	if err != nil {
		return auditlogdto.ListAuditLogsRequest{}, err
	}

	return auditlogdto.ListAuditLogsRequest{
		OrganizationIDs: organizationIDs,
		DateFrom:        dateFrom,
		DateTo:          dateTo,
		Modules:         parseStringList(values["modules"]),
		Actions:         parseStringList(values["actions"]),
		Usernames:       parseStringList(values["usernames"]),
		EntityIDs:       entityIDs,
		EntityTypes:     parseStringList(values["entity_types"]),
		Search:          strings.TrimSpace(values.Get("search")),
		Page:            page,
		PageSize:        pageSize,
		SortBy:          query.SortBy,
		SortOrder:       query.SortOrder,
	}, nil
}

func parseOptionalInt(raw string, field string) (int, error) {
	clean := strings.TrimSpace(raw)
	if clean == "" {
		return 0, nil
	}

	parsed, err := strconv.Atoi(clean)
	if err != nil {
		return 0, newQueryValidationError(field, "number", fmt.Sprintf("%s must be a valid integer", field))
	}

	return parsed, nil
}

func parseRFC3339(raw string, field string) (time.Time, error) {
	clean := strings.TrimSpace(raw)

	parsed, err := time.Parse(time.RFC3339Nano, clean)
	if err == nil {
		return parsed.UTC(), nil
	}

	parsed, err = time.Parse(time.RFC3339, clean)
	if err == nil {
		return parsed.UTC(), nil
	}

	return time.Time{}, newQueryValidationError(field, "rfc3339", fmt.Sprintf("%s must be RFC3339", field))
}

func parseInt64List(values []string, field string) ([]int64, error) {
	items := splitValues(values)
	result := make([]int64, 0, len(items))

	for _, item := range items {
		parsed, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, newQueryValidationError(field, "number", fmt.Sprintf("%s must contain valid integers", field))
		}
		result = append(result, parsed)
	}

	return result, nil
}

func parseStringList(values []string) []string {
	return splitValues(values)
}

func splitValues(values []string) []string {
	result := make([]string, 0, len(values))

	for _, value := range values {
		parts := strings.Split(value, ",")
		for _, part := range parts {
			clean := strings.TrimSpace(part)
			if clean == "" {
				continue
			}
			result = append(result, clean)
		}
	}

	return result
}

func newQueryValidationError(field string, rule string, message string) error {
	return kernel.NewValidationErrorWithData(
		message,
		map[string]any{
			"violations": []sharedvalidator.FieldViolation{
				{Field: field, Rule: rule},
			},
		},
	)
}
