package dto

type LoginResponse struct {
	Token    string `json:"token,omitempty"`
	Username string
}
