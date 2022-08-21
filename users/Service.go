package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lambsroad/go-shop/users/dto"
	"golang.org/x/crypto/bcrypt"
)

var (
	DefaultSecret      = []byte("secret")
	ErrInvalidPassword = errors.New("invalid password")
)

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateUser(request dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	bytes, passwordHashErr := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if passwordHashErr != nil {
		return dto.CreateUserResponse{}, passwordHashErr
	}

	newUser := User{uuid.New().String(), request.UserName, string(bytes)}
	save, err := s.repository.Save(&newUser)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{
		ID:       save.UUID,
		UserName: newUser.UserName,
	}, nil
}

func (s Service) Login(name string, password string) (dto.LoginResponse, error) {
	user, err := s.repository.FindByName(name)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if user.Password != password {
		return dto.LoginResponse{}, err
	}
	token, tokenErr := NewUserToken(user.UserName)

	if tokenErr != nil {
		return dto.LoginResponse{}, tokenErr
	}

	refreshToken, refreshTokenErr := NewUserRefreshToken(user.UserName)
	if refreshTokenErr != nil {
		return dto.LoginResponse{}, refreshTokenErr
	}

	return dto.LoginResponse{Token: token.Token, RefreshToken: refreshToken.RefreshToken, Username: name}, nil
}

func (s *Service) FindUserOneByName(name string) (dto.GetUserResponse, error) {
	foundUser, err := s.repository.FindByName(name)
	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		ID:       foundUser.UUID,
		UserName: foundUser.UserName,
	}, nil
}

func (s *Service) FindUserOneByID(id string) (dto.GetUserResponse, error) {
	foundUser, err := s.repository.FindById(id)
	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		ID:       foundUser.UUID,
		UserName: foundUser.UserName,
	}, nil
}
