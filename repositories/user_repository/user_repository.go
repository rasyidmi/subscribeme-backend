package user_repository

import "projects-subscribeme-backend/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	UpdateFcmTokenUser(username string, fcmToken string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByNpm(npm string) (models.User, error)
}
