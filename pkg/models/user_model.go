package models

import (
	"fmt"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"

	"gorm.io/gorm"
)

type UserModel interface {
	CreateUser(user modelsConfig.User) error
	FindByID(id string) (modelsConfig.User, error)
	FindByEmail(email string) (modelsConfig.User, error)
}

type userModel struct {
	DB *gorm.DB
}

func CreateUserModel(DB *gorm.DB) *userModel {
	return &userModel{DB}
}

func (r *userModel) CreateUser(user modelsConfig.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		fmt.Printf("ERROR OCCURED: %s", err)
	}

	return err
}

func (r *userModel) FindByID(id string) (modelsConfig.User, error) {
	var user modelsConfig.User
	err := r.DB.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the user.")
		return user, err
	}
	return user, err
}

func (r *userModel) FindByEmail(email string) (modelsConfig.User, error) {
	var user modelsConfig.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the user.")
		return user, err
	}
	return user, err
}
