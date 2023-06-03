package models

import "time"

type Job struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Payload string
	RunAt   time.Time
	Cron    *string
}
