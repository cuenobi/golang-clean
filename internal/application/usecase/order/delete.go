package usecase

import "context"

func (u *OrderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
