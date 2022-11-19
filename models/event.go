package models

import (
	"time"
)

type Event struct {
	ID           int       `gorm:"primaryKey"`
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	DeadlineDate time.Time `gorm:"not null"`
	Classes      []*Class  `gorm:"many2many:class_events;constraint:OnDelete:CASCADE"`
	Subject      Subject   `gorm:"constraint:OnDelete:CASCADE;"`
	SubjectID    int       `gorm:"not null"`
	SubjectName  string    `gorm:"not null"`
}

type EventDTO struct {
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description"`
	DeadlineDate time.Time `json:"deadline_date" binding:"required"`
	ClassesID    []int     `json:"classes_id"`
	SubjectID    int       `json:"subject_id" binding:"required"`
	SubjectName  string    `json:"subject_name" binding:"required"`
}

type EventResponse struct {
	ID           int
	Title        string
	Description  string
	DeadlineDate time.Time
	SubjectID    int
	SubjectName  string
}
