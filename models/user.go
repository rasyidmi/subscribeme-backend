package models

type User struct {
	ID     string `gorm:"primaryKey"`
	Email  string `gorm:"not null;unique"`
	Name   string `gorm:"not null"`
	Role   string `gorm:"not null"`
	Avatar string
}

type UserRequest struct {
	ID     string `json:"id" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Role   string `json:"role" binding:"required"`
	Avatar string `json:"avatar"`
}
