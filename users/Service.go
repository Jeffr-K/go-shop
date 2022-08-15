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
	Repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) CreateUser(request dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	newUser := User{uuid.New().String(), request.UserName, request.Password}
	save, err := s.Repository.Save(&newUser)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{
		ID:       save.UUID,
		UserName: newUser.UserName,
	}, nil
}

func (s Service) Login(name string, password string) (dto.LoginResponse, error) {
	user, err := s.Repository.FindByName(name)
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

func (s *Service) FindUserOneByName(name string) (dto.GetUserResponse, error) {
	foundUser, err := s.Repository.FindByName(name)
	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		ID:       foundUser.UUID,
		UserName: foundUser.UserName,
	}, nil
}

func (s *Service) FindUserOneByID(id string) (dto.GetUserResponse, error) {
	foundUser, err := s.Repository.FindById(id)
	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		ID:       foundUser.UUID,
		UserName: foundUser.UserName,
	}, nil
}
