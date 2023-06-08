package course_repository

import (
	"projects-subscribeme-backend/models"
)

type CourseRepository interface {
	CreateCourse(course models.CourseScele, user models.User) (models.CourseScele, error)
	FirstOrCreateEvent(event models.ClassEvent) (models.ClassEvent, error)
	DeleteUserEventByUserIdAndCourseId(userId string, courseId string) error
	DeletUserCourseByUserAndCourse(user models.User, course models.CourseScele) error
	CreateUserEvent(userEvent models.UserEvent) (models.UserEvent, error)
	GetEventByCourseId(courseId string) ([]models.ClassEvent, error)
	GetEventByEventId(eventId string) (models.ClassEvent, error)
	GetCourseByCourseID(courseId string) (models.CourseScele, error)
	GetCourseByCourseSceleId(courseId int64) (models.CourseScele, error)
	GetUserEventByCourseId(courseId string, userId string) ([]models.UserEvent, error)

	GetUserCourseByUsername(username string) ([]*models.CourseScele, error)
	GetDeadlineTodayByUserId(userId string) ([]models.UserEvent, error)
	GetDeadline7DaysAheadByUserId(userId string) ([]models.UserEvent, error)

	DeleteJobsByUserIdAndCourseId(userId string, courseId string) error
}
