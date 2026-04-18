package persistence

import (
	"context"
	"errors"

	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	"github.com/cuenobi/golang-clean/internal/user/application/port/out"
	"github.com/cuenobi/golang-clean/internal/user/domain/entity"
	"github.com/cuenobi/golang-clean/internal/user/domain/valueobject"
	"gorm.io/gorm"
)

var _ out.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	model := UserModel{ID: user.ID, Name: user.Name, Email: string(user.Email), CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	return sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Create(&model).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var model UserModel
	err := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).First(&model, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kernel.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toEntity(model)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var model UserModel
	err := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).First(&model, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kernel.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toEntity(model)
}

func (r *UserRepository) List(ctx context.Context) ([]*entity.User, error) {
	var models []UserModel
	if err := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Order("created_at desc").Find(&models).Error; err != nil {
		return nil, err
	}
	result := make([]*entity.User, 0, len(models))
	for _, model := range models {
		user, err := toEntity(model)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	updates := map[string]any{"name": user.Name, "email": string(user.Email), "updated_at": user.UpdatedAt}
	res := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Model(&UserModel{}).Where("id = ?", user.ID).Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return kernel.ErrNotFound
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	res := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Delete(&UserModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return kernel.ErrNotFound
	}
	return nil
}

func toEntity(model UserModel) (*entity.User, error) {
	email, err := valueobject.NewEmail(model.Email)
	if err != nil {
		return nil, err
	}
	return &entity.User{ID: model.ID, Name: model.Name, Email: email, CreatedAt: model.CreatedAt, UpdatedAt: model.UpdatedAt}, nil
}
