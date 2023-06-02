package course_repository

import (
	"log"
	"projects-subscribeme-backend/models"
	"time"

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
	return courseExists, err

}

func (r *courseRepository) FirstOrCreateEvent(event models.ClassEvent) (models.ClassEvent, error) {
	err := r.db.FirstOrCreate(&event, models.ClassEvent{CourseSceleID: event.CourseSceleID, EventName: event.EventName, Date: event.Date}).Error

	return event, err
}

func (r *courseRepository) CreateUserEvent(userEvent models.UserEvent) (models.UserEvent, error) {
	err := r.db.Create(&userEvent).Error

	return userEvent, err
}

func (r *courseRepository) GetEventByCourseId(courseId string) ([]models.ClassEvent, error) {
	var event []models.ClassEvent

	err := r.db.Find(&event, "course_scele_id = ?", courseId).Error

	return event, err
}

func (r *courseRepository) GetUserEventByCourseId(courseId string, userId string) ([]models.UserEvent, error) {
	var userEvents []models.UserEvent

	err := r.db.Preload("ClassEvent").Find(&userEvents, "user_id = ? AND course_id = ?", userId, courseId).Error

	return userEvents, err

}

func (r *courseRepository) GetUserCourseByUsername(username string) ([]*models.CourseScele, error) {
	var user models.User

	err := r.db.Preload("CourseScele").First(&user, "username = ?", username).Error

	return user.CourseScele, err

}

func (r *courseRepository) GetDeadlineTodayByUserId(userId string) ([]models.UserEvent, error) {
	var userEvents []models.UserEvent

	currentDate := time.Now().Format("2006-01-02")

	err := r.db.Preload("ClassEvent").Joins("JOIN class_events ce ON ce.id = user_events.event_id").
		Where("user_id = ?", userId).Where("DATE(ce.date) = ?", currentDate).Order("ce.date DESC").Find(&userEvents).Error

	return userEvents, err
}

func (r *courseRepository) GetDeadline7DaysAheadByUserId(userId string) ([]models.UserEvent, error) {
	var userEvents []models.UserEvent

	currentDate := time.Now()
	next7Days := currentDate.AddDate(0, 0, 7)

	err := r.db.Preload("ClassEvent").Joins("JOIN class_events ce ON ce.id = user_events.event_id").
		Where("user_id = ?", userId).Where("DATE(ce.date) BETWEEN ? AND ?", currentDate, next7Days).Order("ce.date DESC").Find(&userEvents).Error

	return userEvents, err

}
