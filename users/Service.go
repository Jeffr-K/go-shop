package user

import (
	"errors"
)

var (
	DefaultSecret      = []byte("secret")
	ErrInvalidPassword = errors.New("invalid password")
)

type Service struct {
	secret []byte
}

func NewService(secret []byte) *Service {
	return &Service{secret: secret}
}

func (s Service) Login(name, password string) {
}
