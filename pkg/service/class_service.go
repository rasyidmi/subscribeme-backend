package service

import (
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/models"
)

type ClassService interface {
	GetClassByID(id int) (modelsConfig.ClassResponse, error)
}

type classService struct {
	repo models.ClassRepository
}

func CreateClassService(repo models.ClassRepository) *classService {
	return &classService{repo}
}

func (s *classService) GetClassByID(id int) (modelsConfig.ClassResponse, error) {
	class, err := s.repo.GetClassByID(id)
	return class, err
}
