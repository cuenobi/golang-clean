package usecase

import (
	"github.com/cuenobi/golang-clean/internal/application/port/in"
	"github.com/cuenobi/golang-clean/internal/application/port/out"
)

var _ in.UserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	repo  out.UserRepository
	clock out.Clock
	idGen out.IDGenerator
}

func NewUserUseCase(repo out.UserRepository, clock out.Clock, idGen out.IDGenerator) *UserUseCase {
	return &UserUseCase{repo: repo, clock: clock, idGen: idGen}
}
