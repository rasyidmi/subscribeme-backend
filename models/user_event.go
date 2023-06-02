package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEvent struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	UserID     string    `gorm:"not null"`
	EventID    int       `gorm:"not null"`
	CourseName string    `gorm:"not null"`
	EventName  string    `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
}

func (userEvent *UserEvent) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	userEvent.ID = uuid.New()
	return
}
