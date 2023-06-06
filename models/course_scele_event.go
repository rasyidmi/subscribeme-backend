package models

import (
	"projects-subscribeme-backend/constant"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassEvent struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	CourseSceleID  string
	CourseModuleID int64 `gorm:"unique"`
	Type           constant.EventEnum
	Date           time.Time
	EventName      string
	CourseScele    *CourseScele
}

func (classEvent *ClassEvent) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	classEvent.ID = uuid.New()
	return
}
