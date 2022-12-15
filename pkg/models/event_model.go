package models

import (
	"fmt"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event modelsConfig.Event, classsesID []int) error
	DeleteByID(id int) error
	FindByID(id int) (modelsConfig.EventResponse, error)
	UpdateByID(id int, newData map[string]interface{}) error
	GetTodayDeadline(userId string) ([]modelsConfig.StudentEvent, error)
}

type eventRepository struct {
	DB *gorm.DB
}

func CreateEventRepository(DB *gorm.DB) *eventRepository {
	return &eventRepository{DB}
}

func (r *eventRepository) Create(event modelsConfig.Event, classsesID []int) error {
	err := r.DB.Debug().Create(&event).Error
	// var err error
	fmt.Println(event)
	for _, classID := range classsesID {
		var class modelsConfig.Class
		err = r.DB.Debug().Table("classes").Where("ID = ?", classID).Find(&class).Error
		fmt.Println(class)
		if err != nil {
			fmt.Printf("ERROR OCCURED: %s", err)
			break
		} else {
			err = r.DB.Debug().Model(&event).Association("Classes").Append(&class)
			if err != nil {
				fmt.Printf("ERROR OCCURED: %s", err)
			}
		}
	}
	return err
}

func (r *eventRepository) FindByID(id int) (modelsConfig.EventResponse, error) {
	var event modelsConfig.EventResponse
	err := r.DB.Debug().Model(&modelsConfig.Event{}).First(&event, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the subject.")
		return event, err
	}

	return event, err
}

func (r *eventRepository) DeleteByID(id int) error {
	err := r.DB.Delete(&modelsConfig.Event{}, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when deleting the event.")
	}
	return err
}

func (r *eventRepository) UpdateByID(id int, newData map[string]interface{}) error {
	event, err := r.FindByID(id)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the subject.")
		return err
	} else if event.SubjectID != newData["subject_id"] && event.SubjectName != newData["subject_name"] {
		fmt.Println("ERROR OCCURED: New data subject_id/name is different with the old one.")
		err = utils.DifferentSubjectID{}
		return err
	}

	err = r.DB.Model(&modelsConfig.Event{}).Where("id = ?", id).Updates(newData).Error
	return err
}

func (r *eventRepository) GetTodayDeadline(userId string) ([]modelsConfig.StudentEvent, error) {
	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
		0, 0, 0, 0, currentTime.Location())
	endOfDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
		24, 0, 0, 0, currentTime.Location())
	events := []modelsConfig.StudentEvent{}
	err := r.DB.Debug().Model(modelsConfig.StudentEvent{}).Where("user_id = ?", userId).
		Where("deadline_date < ?", endOfDate).Where("deadline_date >= ?", currentDate).Find(&events).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the events.")
		return nil, err
	}
	return events, err
}
