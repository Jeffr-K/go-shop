package boards

type IBoardRepository interface {
	FindAll() ([]*Boards, error)
	FindById(id int) (*Boards, error)
	DeleteById(id int) error
	Save(user *Boards) (*Boards, error)
	Update(user *Boards) (*Boards, error)
}
