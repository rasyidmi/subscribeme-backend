package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassAbsenceSession struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	TeacherName string
	ClassCode   string    `gorm:"uniqueIndex:class_session_idx"`
	StartTime   time.Time `gorm:"uniqueIndex:class_session_idx"`
	EndTime     time.Time
	IsGeofence  bool
	GeoRadius   float64
	Latitude    float64   `gorm:"default:null"`
	Longitude   float64   `gorm:"default:null"`
	Absence     []Absence `gorm:"foreignKey:ClassAbsenceSessionID"`
}

type Absence struct {
	ClassAbsenceSessionID string
	StudentName           string
	StudentNpm            string  `gorm:"unique"`
	Latitude              float64 `gorm:"default:null"`
	Longitude             float64 `gorm:"default:null"`
	DeviceCode            string
	PresentTime           time.Time
	Present               bool
}

func (cas *ClassAbsenceSession) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	cas.ID = uuid.New()
	return
}