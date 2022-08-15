package boards

type Boards struct {
	Title   string `gorm:"not_null"`
	Content string `gorm:"not_null"`
	UserID  string `gorm:"not_null"`
}
