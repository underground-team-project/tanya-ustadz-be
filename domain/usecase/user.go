package usecase

import (
	"context"
	"tanyaustadz/domain/entity"
	"tanyaustadz/pkg/exceptions"
)

type UserUseCase interface {
	Login(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomError)
	Register(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomError)
}
