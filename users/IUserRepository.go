package user

type IUserRepository interface {
	FindAll() ([]*User, error)
	FindById(id int) (*User, error)
	DeleteById(id int) error
	Save(user *User) (*User, error)
	Update(user *User) (*User, error)
}
