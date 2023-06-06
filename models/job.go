package models

import (
	"time"
)

type Job struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  string `gorm:"default:null"`
	EventID string `gorm:"default:null"`
	Name    string
	Payload string
	RunAt   time.Time
	Cron    *string
	User    *User
}
