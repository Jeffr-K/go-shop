package user

import (
	"errors"
	"gorm.io/gorm"
)

var ErrNotFoundUser = errors.New("not found user")

type UserRepository struct {
	Repository *gorm.DB
}

func (UserRepository) NewUserRepositoryImpl(Repository *gorm.DB) *UserRepository {
	return &UserRepository{Repository}
}

func (userRepository *UserRepository) FindAll() ([]*User, error) {
	var users []*User
	err := userRepository.Repository.Preload("Cards").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepository *UserRepository) FindById(id string) (*User, error) {
	var user = new(User)
	err := userRepository.Repository.Preload("Boards").First(&user, id).Error

	if err != nil {
		return nil, ErrNotFoundUser
	}
	return user, nil
}

func (userRepository *UserRepository) FindByName(userName string) (*User, error) {
	var user = new(User)
	err := userRepository.Repository.Preload("Boards").First(&user, userName).Error

	if err != nil {
		return nil, ErrNotFoundUser
	}
	return user, nil
}

func (userRepository *UserRepository) DeleteById(id int) error {

	userRepository.Repository.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&User{}, id).Error
		if err != nil {
			return err
		}

		return nil
	})
	return nil
}

func (userRepository *UserRepository) Save(user *User) (*User, error) {
	var userCheck = new(User)
	err1 := userRepository.Repository.Where("name = ?", user.UserName).First(&userCheck).Error
	if err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			err2 := userRepository.Repository.Create(user).Error
			if err2 != nil {
				return nil, err2
			}
			return user, nil
		}
	}
	return nil, errors.New("duplicate user")
}
