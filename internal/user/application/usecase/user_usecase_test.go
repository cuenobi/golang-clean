package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/cuenobi/golang-clean/internal/user/application/dto"
	"github.com/cuenobi/golang-clean/internal/user/application/usecase"
	"github.com/cuenobi/golang-clean/internal/user/domain/entity"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userUseCaseSuite struct {
	suite.Suite
}

func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(userUseCaseSuite))
}

func (s *userUseCaseSuite) TestCRUDFlow() {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	repo := newUserRepoMock()
	uc := usecase.NewUserUseCase(repo, fixedClock{now: now}, fixedID{id: "usr_1"})

	created, err := uc.CreateUser(context.Background(), dto.CreateUserRequest{Name: "Alice", Email: "alice@example.com"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), "usr_1", created.ID)

	got, err := uc.GetUser(context.Background(), "usr_1")
	require.NoError(s.T(), err)
	require.Equal(s.T(), "Alice", got.Name)

	updated, err := uc.UpdateUser(context.Background(), "usr_1", dto.UpdateUserRequest{Name: "Alice A", Email: "alice.a@example.com"})
	require.NoError(s.T(), err)
	require.Equal(s.T(), "Alice A", updated.Name)

	list, err := uc.ListUsers(context.Background())
	require.NoError(s.T(), err)
	require.Len(s.T(), list, 1)

	err = uc.DeleteUser(context.Background(), "usr_1")
	require.NoError(s.T(), err)

	_, err = uc.GetUser(context.Background(), "usr_1")
	require.Error(s.T(), err)
}

type userRepoMock struct {
	users map[string]*entity.User
}

func newUserRepoMock() *userRepoMock {
	return &userRepoMock{users: map[string]*entity.User{}}
}

func (m *userRepoMock) Create(ctx context.Context, user *entity.User) error {
	m.users[user.ID] = user
	return nil
}

func (m *userRepoMock) GetByID(ctx context.Context, id string) (*entity.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, kernel.ErrNotFound
	}
	return u, nil
}

func (m *userRepoMock) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	for _, u := range m.users {
		if string(u.Email) == email {
			return u, nil
		}
	}
	return nil, kernel.ErrNotFound
}

func (m *userRepoMock) List(ctx context.Context) ([]*entity.User, error) {
	result := make([]*entity.User, 0, len(m.users))
	for _, u := range m.users {
		result = append(result, u)
	}
	return result, nil
}

func (m *userRepoMock) Update(ctx context.Context, user *entity.User) error {
	if _, ok := m.users[user.ID]; !ok {
		return errors.New("not found")
	}
	m.users[user.ID] = user
	return nil
}

func (m *userRepoMock) Delete(ctx context.Context, id string) error {
	if _, ok := m.users[id]; !ok {
		return kernel.ErrNotFound
	}
	delete(m.users, id)
	return nil
}

type fixedClock struct {
	now time.Time
}

func (f fixedClock) Now() time.Time { return f.now }

type fixedID struct {
	id string
}

func (f fixedID) NewID() string { return f.id }
