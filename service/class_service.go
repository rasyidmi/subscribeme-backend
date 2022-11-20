package service

import (
	"projects-subscribeme-backend/config/models"
	"projects-subscribeme-backend/repository"
)

type ClassService interface {
	GetClassByID(id int) (models.ClassResponse, error)
}

type classService struct {
	repo repository.ClassRepository
}

func CreateClassService(repo repository.ClassRepository) *classService {
	return &classService{repo}
}

func (s *classService) GetClassByID(id int) (models.ClassResponse, error) {
	class, err := s.repo.GetClassByID(id)
	return class, err
}
