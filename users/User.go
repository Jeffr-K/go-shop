package user

type User struct {
	UUID     string `gorm:"primaryKey"`
	UserName string `gorm:"not_null"`
	Password string `gorm:"not_null"`
}
