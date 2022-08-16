package dto

type CreateBoardRequest struct {
	Title   string `json:"title" example:"board title"`
	Content string `json:"content"  example:"board content"`
}

type UpdateBoardRequest struct {
	BoardId int    `json:"board_id" example:"board_id"`
	Title   string `json:"update board title"`
	Content string `json:"update board content"`
}
