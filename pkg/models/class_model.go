package models

import (
	"fmt"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"

	"gorm.io/gorm"
)

type ClassModel interface {
	GetClassByID(id int, userId string) (modelsConfig.ClassResponse, error)
	Subscribe(id int, userId string) error
}

type classModel struct {
	DB *gorm.DB
}

func CreateClassRepository(DB *gorm.DB) *classModel {
	return &classModel{DB}
}

func (r *classModel) GetClassByID(id int, userId string) (modelsConfig.ClassResponse, error) {
	var class modelsConfig.Class
	var classResponse modelsConfig.ClassResponse
	err := r.DB.Debug().Preload("Events").Preload("Users", "id = ?", userId).First(&class, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the association.")
		return classResponse, err
	}
	// Check if the user is subscribing the class.
	isSubscribe := false
	if len(class.Users) != 0 {
		isSubscribe = true
	}
	// Convert Class object to ClassResponse
	classResponse = modelsConfig.ClassResponse{
		ID:          class.ID,
		Title:       class.Title,
		Events:      class.Events,
		IsSubscribe: isSubscribe,
	}

	return classResponse, err
}

func (r *classModel) Subscribe(id int, userId string) error {
	var user modelsConfig.User
	var class modelsConfig.Class
	// Load the user.
	err := r.DB.Debug().Where("id = ?", userId).First(&user).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR OCCURED: Error when finding the user.")
	}
	// Load the class.
	err = r.DB.Debug().Preload("Events").First(&class, id).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR OCCURED: Error when finding the class.")
	}
	// Connect student with class's events.
	for i := 0; i < len(class.Events); i++ {
		studentEvent := modelsConfig.StudentEvent{
			UserID:       user.ID,
			EventID:      class.Events[i].ID,
			ClassName:    class.Title,
			EventName:    class.Events[i].Title,
			SubjectName:  class.Events[i].SubjectName,
			DeadlineDate: class.Events[i].DeadlineDate,
		}
		err = r.DB.Debug().Model(modelsConfig.StudentEvent{}).Create(&studentEvent).Error
		if err != nil {
			fmt.Println(err)
			fmt.Println("ERROR OCCURED: when adding class event.")
		}
	}
	err = r.DB.Debug().Model(&class).Association("Users").Append(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR OCCURED: Error when subscribing.")
	}

	return err
}
