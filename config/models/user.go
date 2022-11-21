package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    string    `gorm:"not null;unique"`
	Password string    `gorm:"not null"`
	Name     string    `gorm:"not null"`
	Role     string    `gorm:"not null"`
	Avatar   string
}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Avatar   string `json:"avatar"`
}
