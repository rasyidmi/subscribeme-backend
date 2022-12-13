package models

import (
	"fmt"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"

	"gorm.io/gorm"
)

type ClassModel interface {
	GetClassByID(id int) (modelsConfig.ClassResponse, error)
}

type classModel struct {
	DB *gorm.DB
}

func CreateClassRepository(DB *gorm.DB) *classModel {
	return &classModel{DB}
}

func (r *classModel) GetClassByID(id int) (modelsConfig.ClassResponse, error) {
	var class modelsConfig.Class
	var classResponse modelsConfig.ClassResponse
	err := r.DB.Debug().Model(&modelsConfig.Class{}).Find(&class, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the class.")
		return classResponse, err
	}

	// Finding the class's events
	var events []*modelsConfig.EventResponse
	err = r.DB.Debug().Model(&class).Where("class_id = ?", id).Association("Events").Find(&events)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the class's events.")
		return classResponse, err
	}

	// Convert Class object to ClassResponse
	classResponse = modelsConfig.ClassResponse{
		ID:     class.ID,
		Title:  class.Title,
		Events: events,
	}

	return classResponse, err
}
