package course_repository

import (
	"projects-subscribeme-backend/models"
)

type CourseRepository interface {
	CreateCourse(course models.CourseScele, user models.User) (models.CourseScele, error)
	FirstOrCreateEvent(event models.ClassEvent) (models.ClassEvent, error)
	CreateUserEvent(userEvent models.UserEvent) (models.UserEvent, error)
	GetEventByCourseId(courseId string) ([]models.ClassEvent, error)
	GetUserEventByCourseId(courseId string, userId string) ([]models.UserEvent, error)
	GetUserCourseByUsername(username string) ([]*models.CourseScele, error)
	GetDeadlineTodayByUserId(userId string) ([]models.UserEvent, error)
	GetDeadline7DaysAheadByUserId(userId string) ([]models.UserEvent, error)
}
