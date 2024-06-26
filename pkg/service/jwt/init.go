package jwt

import (
	"os"
	"tanyaustadz/pkg/config"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaim struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")

	fields := []string{"JWT_SECRET"}

	for _, f := range fields {
		err := config.Required(f)
		if err != nil {
			panic(err.Error())
		}
	}

	return secretKey
}
