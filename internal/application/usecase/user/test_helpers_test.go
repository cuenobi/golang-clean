package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	dto "github.com/cuenobi/golang-clean/internal/application/dto/user"
	usecase "github.com/cuenobi/golang-clean/internal/application/usecase/user"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	"github.com/stretchr/testify/require"
)

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

type userFixedClock struct {
	now time.Time
}

func (f userFixedClock) Now() time.Time {
	return f.now
}

type userFixedID struct {
	id string
}

func (f userFixedID) NewID() string {
	return f.id
}

func newUserUseCaseForTest(id string, now time.Time) (*usecase.UserUseCase, *userRepoMock) {
	repo := newUserRepoMock()
	uc := usecase.NewUserUseCase(repo, userFixedClock{now: now}, userFixedID{id: id})
	return uc, repo
}

func seedUser(t *testing.T, repo *userRepoMock, id, name, email string, now time.Time) {
	t.Helper()

	voEmail, err := valueobject.NewEmail(email)
	require.NoError(t, err)
	user, err := entity.NewUser(id, name, voEmail, now)
	require.NoError(t, err)
	require.NoError(t, repo.Create(context.Background(), user))
}

func createUserReq() dto.CreateUserRequest {
	return dto.CreateUserRequest{Name: "Alice", Email: "alice@example.com"}
}
