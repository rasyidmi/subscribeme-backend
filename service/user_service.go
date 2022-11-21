package service

import (
	"projects-subscribeme-backend/config/models"
	"projects-subscribeme-backend/repository"
	"projects-subscribeme-backend/utils"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(subject models.UserRequest) error
	FindByID(id string) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func CreateUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (s *userService) CreateUser(user models.UserRequest) error {
	userConverted := models.User{
		Password: user.Password,
		Role:     user.Role,
		Email:    user.Email,
		Name:     user.Name,
		Avatar:   user.Avatar,
	}

	err := s.repo.CreateUser(userConverted)

	return err
}

func (s *userService) FindByID(id string) (models.User, error) {
	user, err := s.repo.FindByID(id)
	if user.ID == uuid.Nil {
		err = utils.DataNotFound{}
	}

	return user, err
}
