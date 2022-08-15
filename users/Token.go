package user

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type ResponseTokenDto struct {
	UserName string `json:"user_name" example:"ash"`
	Token    jwt.StandardClaims
}

func CreateToken(name string) ResponseTokenDto {
	return ResponseTokenDto{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func NewUserToken(name string) (ResponseTokenDto, error) {
	return CreateToken(name), nil
}
