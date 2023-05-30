package user_repository

import "projects-subscribeme-backend/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

}
