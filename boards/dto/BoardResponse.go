package dto

type CreateBoardResponse struct {
	ID      string `json:"id" example:"1"`
	Title   string `json:"title" example:"board title"`
	Content string `json:"content" example:"board_content"`
}

type UpdateBoardResponse struct {
	Title   string `json:"title" example:"update board title"`
	Content string `json:"update board content"`
}

type GetBoardResponse struct {
	Title   string `json:"title" example:"3"`
	Content string `json:"content" example:"content test"`
}

type Fail400GetResponse struct {
	Message string `json:"message" example:"Bad Request"`
}
