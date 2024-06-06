package user

import (
	"tanyaustadz/domain/repository"
	"tanyaustadz/domain/usecase"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(
	userRepository repository.UserRepository,
) usecase.UserUseCase {
	return &userInteractor{
		userRepository: userRepository,
	}
}
