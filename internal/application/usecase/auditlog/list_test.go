package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	auditlogdto "github.com/cuenobi/golang-clean/internal/application/dto/auditlog"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
	usecase "github.com/cuenobi/golang-clean/internal/application/usecase/auditlog"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type auditLogRepoMock struct {
	listFn func(ctx context.Context, filter out.AuditLogListFilter) ([]*entity.AuditLog, int64, error)
}

func (m *auditLogRepoMock) List(ctx context.Context, filter out.AuditLogListFilter) ([]*entity.AuditLog, int64, error) {
	if m.listFn == nil {
		return []*entity.AuditLog{}, 0, nil
	}
	return m.listFn(ctx, filter)
}

type AuditLogUseCaseSuite struct {
	suite.Suite

	ctx  context.Context
	repo *auditLogRepoMock
	uc   *usecase.AuditLogUseCase
}

func TestAuditLogUseCaseSuite(t *testing.T) {
	suite.Run(t, new(AuditLogUseCaseSuite))
}

func (s *AuditLogUseCaseSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = &auditLogRepoMock{}
	s.uc = usecase.NewAuditLogUseCase(s.repo)
}

func (s *AuditLogUseCaseSuite) TestGetSystemAuditLogs_Success() {
	now := time.Date(2026, 4, 25, 9, 0, 0, 0, time.UTC)
	s.repo.listFn = func(ctx context.Context, filter out.AuditLogListFilter) ([]*entity.AuditLog, int64, error) {
		s.Equal(out.AuditLogScopeSystem, filter.Scope)
		s.Equal("SYSTEM_", filter.ModulePrefix)
		s.Equal([]string{"SYSTEM_USER"}, filter.Modules)
		s.Equal(1, filter.Page)
		s.Equal(20, filter.PageSize)

		return []*entity.AuditLog{
			{
				ID:         1001,
				EventID:    "evt-1",
				EntityType: "USER",
				EntityID:   7,
				Action:     "UPDATE",
				Username:   "system",
				Module:     "SYSTEM_USER",
				DiffValue:  []byte(`{"status":"updated"}`),
				OccurredAt: now,
				CreatedAt:  now,
			},
		}, 1, nil
	}

	resp, err := s.uc.GetSystemAuditLogs(s.ctx, auditlogdto.ListAuditLogsRequest{
		DateFrom: now.Add(-1 * time.Hour),
		DateTo:   now,
		Modules:  []string{"SYSTEM_USER", "ORG_ORDER"},
	})

	s.Require().NoError(err)
	s.Equal(int64(1), resp.Total)
	s.Equal(1, resp.Page)
	s.Equal(20, resp.PageSize)
	s.Equal(1, resp.TotalPages)
	s.Require().Len(resp.Data, 1)
	s.Equal("SYSTEM_USER", resp.Data[0].Module)

	diffMap, ok := resp.Data[0].DiffValue.(map[string]any)
	s.Require().True(ok)
	s.Equal("updated", diffMap["status"])
}

func (s *AuditLogUseCaseSuite) TestGetOrganizationAuditLogs_Success() {
	now := time.Date(2026, 4, 25, 9, 0, 0, 0, time.UTC)
	s.repo.listFn = func(ctx context.Context, filter out.AuditLogListFilter) ([]*entity.AuditLog, int64, error) {
		s.Equal(out.AuditLogScopeOrganization, filter.Scope)
		s.Equal("ORG_", filter.ModulePrefix)
		s.Equal([]string{"ORG_ORDER"}, filter.Modules)
		s.Equal([]int64{1711}, filter.OrganizationIDs)
		s.Equal(2, filter.Page)
		s.Equal(10, filter.PageSize)
		return []*entity.AuditLog{}, 0, nil
	}

	resp, err := s.uc.GetOrganizationAuditLogs(s.ctx, auditlogdto.ListAuditLogsRequest{
		DateFrom:        now.Add(-2 * time.Hour),
		DateTo:          now,
		OrganizationIDs: []int64{1711},
		Modules:         []string{"ORG_ORDER", "SYSTEM_USER"},
		Page:            2,
		PageSize:        10,
	})

	s.Require().NoError(err)
	s.Equal(int64(0), resp.Total)
	s.Equal(2, resp.Page)
	s.Equal(10, resp.PageSize)
	s.Equal(0, resp.TotalPages)
	s.Len(resp.Data, 0)
}

func (s *AuditLogUseCaseSuite) TestGetSystemAuditLogs_RepositoryError() {
	expectedErr := errors.New("db unavailable")
	s.repo.listFn = func(ctx context.Context, filter out.AuditLogListFilter) ([]*entity.AuditLog, int64, error) {
		return nil, 0, expectedErr
	}

	_, err := s.uc.GetSystemAuditLogs(s.ctx, auditlogdto.ListAuditLogsRequest{
		DateFrom: time.Date(2026, 4, 25, 8, 0, 0, 0, time.UTC),
		DateTo:   time.Date(2026, 4, 25, 9, 0, 0, 0, time.UTC),
	})

	require.Error(s.T(), err)
	require.ErrorIs(s.T(), err, expectedErr)
}
