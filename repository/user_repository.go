package repository

import (
	"fmt"
	"projects-subscribeme-backend/config/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	FindByID(id string) (models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB}
}

func (r *userRepository) CreateUser(user models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		fmt.Printf("ERROR OCCURED: %s", err)
	}

	return err
}

func (r *userRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.DB.Debug().Where("id = ?", id).Find(&user).Error
	if err != nil {
		fmt.Println("ERROR OCCURED: Error when finding the user.")
		return user, err
	}
	return user, err
}
