package usecase

import "context"

func (u *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
