package user

type User struct {
	ID       string
	UserName string
	Password string
}

func (u User) CreateUser(user User) User {
	newUser := User{
		ID:       user.ID,
		UserName: user.UserName,
		Password: user.Password,
	}
	return newUser
}

func (u User) DeleteUser() {}

func (u User) UpdateUser() {}

func (u User) SearchUser() {}
