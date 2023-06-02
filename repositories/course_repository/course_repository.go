package course_repository

import (
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"
)

type CourseRepository interface {
	CreateCourse(course models.CourseScele, user models.User) (models.CourseScele, error)
	FirstOrCreateEvent(event models.ClassEvent, eventType constant.EventEnum) (models.ClassEvent, error)
}
