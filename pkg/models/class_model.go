package models

import (
	"fmt"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"

	"gorm.io/gorm"
)

type ClassModel interface {
	GetClassByID(id int) (modelsConfig.ClassResponse, error)
	Subscribe(id int, userId string) error
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

func (r *classModel) Subscribe(id int, userId string) error {
	var user modelsConfig.User
	var class modelsConfig.Class
	err := r.DB.Debug().Where("id = ?", userId).First(&user).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR OCCURED: Error when finding the user.")
	}
	err = r.DB.Debug().Where("id = ?", id).First(&class).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR OCCURED: Error when finding the class.")
	}

	err = r.DB.Debug().Model(&class).Association("Users").Append(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR OCCURED: Error when subscribing.")
	}
	return err
}
