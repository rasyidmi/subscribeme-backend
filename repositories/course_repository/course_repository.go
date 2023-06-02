package course_repository

import "projects-subscribeme-backend/models"

type CourseRepository interface {
	FirstOrCreateCourse(course models.CourseScele) (models.CourseScele, error)
	FirstOrCreateEvent(event models.ClassEvent) (models.ClassEvent, error)
}
