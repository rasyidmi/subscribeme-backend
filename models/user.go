package models

import (
	"projects-subscribeme-backend/constant"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Username    string    `gorm:"unique"`
	Role        constant.UserRoleEnum
	FcmToken    string         `gorm:"fcm_token"`
	CourseScele []*CourseScele `gorm:"many2many:user_course;"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}
