package repository

import (
	"fmt"
	"projects-subscribeme-backend/config/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	GetClassByID(id int) (models.ClassResponse, error)
}

type classRepository struct {
	DB *gorm.DB
}

func CreateClassRepository(DB *gorm.DB) *classRepository {
	return &classRepository{DB}
}

func (r *classRepository) GetClassByID(id int) (models.ClassResponse, error) {
	var class models.Class
	var classResponse models.ClassResponse
	err := r.DB.Debug().Model(&models.Class{}).Find(&class, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the class.")
		return classResponse, err
	}

	// Finding the class's events
	var events []*models.EventResponse
	err = r.DB.Debug().Model(&class).Where("class_id = ?", id).Association("Events").Find(&events)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the class's events.")
		return classResponse, err
	}

	// Convert Class object to ClassResponse
	classResponse = models.ClassResponse{
		ID:     class.ID,
		Title:  class.Title,
		Events: events,
	}

	return classResponse, err
}
