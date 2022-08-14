package user

import (
	"errors"
)

var ErrNotFoundUser = errors.New("not found user")

type Repository struct {
	data map[string]User
}

func NewRepository(data map[string]User) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(user User) {
	r.data[user.ID] = user
}

func (r *Repository) GetByName(userName string) (User, error) {
	for _, user := range r.data {
		if user.UserName == userName {
			return user, nil
		}
	}
	return User{}, ErrNotFoundUser
}

func (r *Repository) GetById(id string) (User, error) {
	for _, user := range r.data {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, ErrNotFoundUser
}

func (r *Repository) Update(user User) {
	r.data[user.ID] = user
}

func (r *Repository) DeleteById(id string) {
	delete(r.data, id)
}
