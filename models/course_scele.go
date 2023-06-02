package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseScele struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	CourseSceleID   int       `gorm:"not null"`
	CourseSceleName string
	ClassEvents     []*ClassEvent `gorm:"many2many:class_events"`
	User            []*User       `gorm:"many2many:user_course;"`
}

func (courseScele *CourseScele) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	courseScele.ID = uuid.New()
	return
}
