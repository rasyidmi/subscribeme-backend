package models

import (
	"projects-subscribeme-backend/constant"
	"time"

	"github.com/google/uuid"
)

type ClassEvent struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	CourseSceleID string
	Type          constant.EventEnum
	Date          time.Time
	EventName     string
	CourseScele   *CourseScele
}
