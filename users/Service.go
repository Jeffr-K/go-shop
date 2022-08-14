package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lambsroad/go-shop/users/dto"
)

var (
	DefaultSecret      = []byte("secret")
	ErrInvalidPassword = errors.New("invalid password")
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateUser(request dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	newUser := User{uuid.New().String(), request.UserName, request.Password}
	s.repository.Create(newUser)
	return dto.CreateUserResponse{
		ID:       newUser.ID,
		UserName: newUser.UserName,
	}, nil
}

func (s Service) Login(name string, password string) (dto.LoginResponse, error) {
	user, err := s.repository.GetByName(name)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if user.Password != password {
		return dto.LoginResponse{}, err
	}
	token, err := NewUserToken(user.UserName)

	if err != nil {
		return dto.LoginResponse{}, err
	}
	return dto.LoginResponse{Token: token.Token, Username: name}, nil
}

func (s *Service) GetUserOneByName(name string) (dto.GetUserResponse, error) {
	foundUser, err := s.repository.GetByName(name)
	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		ID:       foundUser.ID,
		UserName: foundUser.UserName,
	}, nil
}

func (s *Service) GetUserOneByID(id string) (dto.GetUserResponse, error) {
	foundUser, err := s.repository.GetById(id)
	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		ID:       foundUser.ID,
		UserName: foundUser.UserName,
	}, nil
}
