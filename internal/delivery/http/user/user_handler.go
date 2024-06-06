package user_handler

import (
	"tanyaustadz/domain/repository"
	"tanyaustadz/domain/usecase"
	"tanyaustadz/internal/usecase/user"
	"tanyaustadz/pkg/service/jwt"

	"github.com/gorilla/mux"
)

type userHandler struct {
	jwtService  jwt.JWTService
	userUseCase usecase.UserUseCase
}

func UserHandler(
	r *mux.Router,
	jwtService jwt.JWTService,
	userRepository repository.UserRepository,
) {
	userUseCase := user.NewUserInteractor(userRepository)
	handler := &userHandler{
		jwtService:  jwtService,
		userUseCase: userUseCase,
	}
	r.HandleFunc("/apis/v1/user/login", handler.Login).Methods("POST")
	r.HandleFunc("/apis/v1/user/register", handler.Register).Methods("POST")
}
