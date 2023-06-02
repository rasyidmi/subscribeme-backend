package course_repository

import (
	"log"
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

func (r *courseRepository) CreateCourse(course models.CourseScele, user models.User) (models.CourseScele, error) {

	var courseExists models.CourseScele
	check := r.db.Where("course_scele_id = ?", course.CourseSceleID).Limit(1).Find(&courseExists)
	if check.Error != nil {
		return models.CourseScele{}, check.Error
	}

	exists := check.RowsAffected > 0

	if !exists {
		err := r.db.Create(&course).Error
		if err != nil {
			return models.CourseScele{}, err
		}
		courseExists = course
	}

	log.Println(user)

	err := r.db.Debug().Model(&courseExists).Omit("User.*").Association("User").Append(&user)
	if err != nil {
		return models.CourseScele{}, err
	}
	return course, err

}

func (r *courseRepository) FirstOrCreateEvent(event models.ClassEvent, eventType constant.EventEnum) (models.ClassEvent, error) {
	err := r.db.FirstOrCreate(&event).Error

	return event, err
}
