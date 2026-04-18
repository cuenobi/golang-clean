package persistence

import (
	"context"
	"errors"

	"github.com/cuenobi/golang-clean/internal/application/port/out"
	"github.com/cuenobi/golang-clean/internal/domain/entity"
	"github.com/cuenobi/golang-clean/internal/domain/valueobject"
	"github.com/cuenobi/golang-clean/internal/shared/kernel"
	sharedpersistence "github.com/cuenobi/golang-clean/internal/shared/persistence"
	"gorm.io/gorm"
)

var _ out.OrderRepository = (*OrderRepository)(nil)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	model := OrderModel{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Currency:   order.Amount.Currency,
		Amount:     order.Amount.Amount,
		Status:     string(order.Status),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
	return sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Create(&model).Error
}

func (r *OrderRepository) GetByID(ctx context.Context, orderID string) (*entity.Order, error) {
	var model OrderModel
	err := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).First(&model, "id = ?", orderID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kernel.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toOrderEntity(model)
}

func (r *OrderRepository) List(ctx context.Context) ([]*entity.Order, error) {
	var models []OrderModel
	if err := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Order("created_at desc").Find(&models).Error; err != nil {
		return nil, err
	}
	result := make([]*entity.Order, 0, len(models))
	for _, model := range models {
		order, err := toOrderEntity(model)
		if err != nil {
			return nil, err
		}
		result = append(result, order)
	}
	return result, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *entity.Order) error {
	updates := map[string]any{
		"customer_id": order.CustomerID,
		"currency":    order.Amount.Currency,
		"amount":      order.Amount.Amount,
		"status":      string(order.Status),
		"updated_at":  order.UpdatedAt,
	}
	res := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Model(&OrderModel{}).Where("id = ?", order.ID).Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return kernel.ErrNotFound
	}
	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, orderID string) error {
	res := sharedpersistence.FromContext(ctx, r.db).WithContext(ctx).Delete(&OrderModel{}, "id = ?", orderID)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return kernel.ErrNotFound
	}
	return nil
}

func toOrderEntity(model OrderModel) (*entity.Order, error) {
	money, err := valueobject.NewMoney(model.Currency, model.Amount)
	if err != nil {
		return nil, err
	}

	return &entity.Order{
		ID:         model.ID,
		CustomerID: model.CustomerID,
		Amount:     money,
		Status:     entity.Status(model.Status),
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}, nil
}
