package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentEvent struct {
	ID           int       `gorm:"primaryKey"`
	UserID       uuid.UUID `gorm:"not null"`
	EventID      int       `gorm:"not null"`
	SubjectName  string    `gorm:"not null"`
	ClassName    string    `gorm:"not null"`
	EventName    string    `gorm:"not null"`
	DeadlineDate time.Time `gorm:"not null"`
}
