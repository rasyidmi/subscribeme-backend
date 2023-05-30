package user_repository

import (
	"projects-subscribeme-backend/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err

}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username = ?", username).Error

	return user, err
}


