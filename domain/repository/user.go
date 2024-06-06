package repository

import (
	"context"
	"tanyaustadz/domain/entity"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	StoreUser(ctx context.Context, user *entity.User) error
}
