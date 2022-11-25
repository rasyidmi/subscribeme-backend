package models

import (
	"fmt"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	GetAll() ([]modelsConfig.SubjectResponse, error)
	Create(subject modelsConfig.Subject, classes []map[string]string) error
	FindByID(id int) (modelsConfig.Subject, error)
	UpdateByID(id int, newData map[string]interface{}) error
	DeleteByID(id int) error
}

type subjectRepository struct {
	DB *gorm.DB
}

func CreateSubjectRepository(DB *gorm.DB) *subjectRepository {
	return &subjectRepository{DB}
}

func (r *subjectRepository) GetAll() ([]modelsConfig.SubjectResponse, error) {
	var subjects []modelsConfig.SubjectResponse
	err := r.DB.Order("Title").Model(&modelsConfig.Subject{}).Find(&subjects).Error
	if err != nil {
		fmt.Printf("ERROR OCCURED: %s", err)
	}
	return subjects, err
}

func (r *subjectRepository) Create(subject modelsConfig.Subject, classes []map[string]string) error {
	err := r.DB.Create(&subject).Error
	if err != nil {
		fmt.Printf("ERROR OCCURED: %s", err)
	} else {
		for _, class := range classes {
			err = r.DB.Model(&subject).Association("Classes").Append(
				&modelsConfig.Class{
					Title: class["title"],
				},
			)
			if err != nil {
				fmt.Printf("ERROR OCCURED: %s", err)
				break
			}
		}
	}

	return err
}

func (r *subjectRepository) FindByID(id int) (modelsConfig.Subject, error) {
	// Finding the subject
	var subject modelsConfig.Subject
	err := r.DB.Debug().Find(&subject, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the subject.")
		return subject, err
	}

	// Finding the subject's classes
	var classes []modelsConfig.Class
	err = r.DB.Debug().Where("subject_id = ?", id).Find(&classes).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the subject's classes.")
		return subject, err
	}

	subject.Classes = classes

	return subject, err
}

func (r *subjectRepository) UpdateByID(id int, newData map[string]interface{}) error {
	subject, err := r.FindByID(id)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the subject.")
		return err
	}

	err = r.DB.Model(&subject).Updates(newData).Error
	return err
}

func (r *subjectRepository) DeleteByID(id int) error {
	err := r.DB.Debug().Delete(&modelsConfig.Subject{}, id).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when deleting the subject.")
	}
	return err
}
