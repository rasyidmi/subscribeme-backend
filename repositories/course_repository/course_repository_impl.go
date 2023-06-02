package course_repository

import (
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"

	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) FirstOrCreateCourse(course models.CourseScele) (models.CourseScele, error) {
	err := r.db.FirstOrCreate(&course, models.CourseScele{CourseSceleID: course.CourseSceleID}).Error

	return course, err
}

func (r *courseRepository) FirstOrCreateEvent(event models.ClassEvent, eventType constant.EventEnum) (models.ClassEvent, error) {

}
