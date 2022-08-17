package user

type User struct {
	UUID     string `gorm:"primaryKey"`
	UserName string `gorm:"not_null"`
	Password string `gorm:"not_null"`
}

func (u User) CreateUser(user User) User {
	newUser := User{
		UUID:     user.UUID,
		UserName: user.UserName,
		Password: user.Password,
	}
	return newUser
}

func (u User) DeleteUser() {}

func (u User) UpdateUser() {}

func (u User) SearchUser() {}
