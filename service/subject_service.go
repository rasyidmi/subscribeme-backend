package service

import (
	"projects-subscribeme-backend/config/models"
	"projects-subscribeme-backend/repository"
	"projects-subscribeme-backend/utils"
)

type SubjectService interface {
	GetAll() ([]models.SubjectResponse, error)
	Create(subject models.SubjectRequest) error
	FindByID(id int) (models.Subject, error)
	UpdateByID(id int, newData models.SubjectRequest) error
	DeleteByID(id int) error
}

type subjectService struct {
	repo repository.SubjectRepository
}

func CreateSubjectService(repo repository.SubjectRepository) *subjectService {
	return &subjectService{repo}
}

func (s *subjectService) GetAll() ([]models.SubjectResponse, error) {
	subjects, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *subjectService) Create(subjectRequest models.SubjectRequest) error {

	subject := models.Subject{
		Title: subjectRequest.Title,
		Term:  subjectRequest.Term,
		Major: subjectRequest.Major,
	}

	err := s.repo.Create(subject, subjectRequest.Classes)

	return err
}

func (s *subjectService) FindByID(id int) (models.Subject, error) {
	subject, err := s.repo.FindByID(id)
	if subject.ID == 0 {
		err = utils.DataNotFound{}
	}

	return subject, err
}

func (s *subjectService) UpdateByID(id int, newData models.SubjectRequest) error {
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
