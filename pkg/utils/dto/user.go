package dto

import "github.com/google/uuid"

type UserRequest struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
}

type UserResponse struct {
	ID     uuid.UUID `json:"id"`
	Email  string    `json:"email" binding:"email"`
	Name   string    `json:"name"`
	Role   string    `json:"role"`
	Avatar string    `json:"avatar"`
}
