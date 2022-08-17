package user

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type ResponseTokenDto struct {
	UserName string `json:"user_name" example:"ash"`
	Token    jwt.StandardClaims
}

type ResponseRefreshTokenDto struct {
	UserName     string `json:"user_name" example:"ash"`
	RefreshToken jwt.StandardClaims
}

func CreateToken(name string) ResponseTokenDto {
	return ResponseTokenDto{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func CreateRefreshToken(name string) ResponseRefreshTokenDto {
	return ResponseRefreshTokenDto{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
}

func NewUserToken(name string) (ResponseTokenDto, error) {
	return CreateToken(name), nil
}

func NewUserRefreshToken(name string) (ResponseRefreshTokenDto, error) {
	return CreateRefreshToken(name), nil
}
