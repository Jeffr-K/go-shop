package dto

import "github.com/golang-jwt/jwt"

type CreateUserResponse struct {
	ID       string `json:"id" example:"342130dc-f7979-89ec-b939-0232yq12383"`
	UserName string `json:"user_name" example:"andy"`
	Password string `json:"password" example:"#!@@sdflkjsdflkj32"`
}

type LoginResponse struct {
	Token    jwt.StandardClaims `json:"token,omitempty"`
	Username string
}

type UpdateUserResponse struct {
	ID       string `json:"id" example:"342130dc-f7979-89ec-b939-0232yq12383"`
	UserName string `json:"user_name"`
}

type GetUserResponse struct {
	ID       string `json:"id" example:"342130dc-f7979-89ec-b939-0232yq12383"`
	UserName string `json:"user_name" example:"ash"`
}

type Fail400GetResponse struct {
	Message string `json:"message" example:"Bad Request"`
}
