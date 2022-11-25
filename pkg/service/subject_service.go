package service

import (
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/models"
	"projects-subscribeme-backend/pkg/utils"
)

type SubjectService interface {
	GetAll() ([]modelsConfig.SubjectResponse, error)
	Create(subject modelsConfig.SubjectRequest) error
	FindByID(id int) (modelsConfig.Subject, error)
	UpdateByID(id int, newData modelsConfig.SubjectRequest) error
	DeleteByID(id int) error
}

type subjectService struct {
	repo models.SubjectRepository
}

func CreateSubjectService(repo models.SubjectRepository) *subjectService {
	return &subjectService{repo}
}

func (s *subjectService) GetAll() ([]modelsConfig.SubjectResponse, error) {
	subjects, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *subjectService) Create(subjectRequest modelsConfig.SubjectRequest) error {

	subject := modelsConfig.Subject{
		Title: subjectRequest.Title,
		Term:  subjectRequest.Term,
		Major: subjectRequest.Major,
	}

	err := s.repo.Create(subject, subjectRequest.Classes)

	return err
}

func (s *subjectService) FindByID(id int) (modelsConfig.Subject, error) {
	subject, err := s.repo.FindByID(id)
	if subject.ID == 0 {
		err = utils.DataNotFound{}
	}

	return subject, err
}

func (s *subjectService) UpdateByID(id int, newData modelsConfig.SubjectRequest) error {
	data := map[string]interface{}{
		"title": newData.Title,
		"term":  newData.Term,
		"major": newData.Major,
	}

	err := s.repo.UpdateByID(id, data)
	return err
}

func (s *subjectService) DeleteByID(id int) error {
	err := s.repo.DeleteByID(id)
	return err
}
