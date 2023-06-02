package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEvent struct {
	ID         uuid.UUID   `gorm:"primaryKey"`
	UserID     string      `gorm:"not null"`
	EventID    string      `gorm:"not null"`
	CourseID   string      `gorm:"not null"`
	IsDone     bool        `gorm:"not null"`
	User       *User       `gorm:"foreignKey:UserID"`
	ClassEvent *ClassEvent `gorm:"foreignKey:EventID"`
}

func (userEvent *UserEvent) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	userEvent.ID = uuid.New()
	return
}
