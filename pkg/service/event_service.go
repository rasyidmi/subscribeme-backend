package service

import (
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/models"
	"projects-subscribeme-backend/pkg/utils"
)

type EventService interface {
	Create(eventRequest modelsConfig.EventDTO) error
	DeleteByID(id int) error
	FindByID(id int) (modelsConfig.EventResponse, error)
	UpdateByID(id int, newData modelsConfig.EventDTO) error
}

type eventService struct {
	repo models.EventRepository
}

func CreateEventService(repo models.EventRepository) *eventService {
	return &eventService{repo}
}

func (s *eventService) Create(eventRequest modelsConfig.EventDTO) error {
	var classesID []int
	for _, classID := range eventRequest.ClassesID {
		classesID = append(classesID, classID)
	}

	event := modelsConfig.Event{
		Title:        eventRequest.Title,
		Description:  eventRequest.Description,
		DeadlineDate: eventRequest.DeadlineDate,
		SubjectID:    eventRequest.SubjectID,
	}

	err := s.repo.Create(event, classesID)
	return err
}

func (s *eventService) FindByID(id int) (modelsConfig.EventResponse, error) {
	event, err := s.repo.FindByID(id)
	if event.ID == 0 {
		err = utils.DataNotFound{}
	}

	return event, err
}

func (s *eventService) UpdateByID(id int, newData modelsConfig.EventDTO) error {
	data := map[string]interface{}{
		"title":         newData.Title,
		"description":   newData.Description,
		"deadline_date": newData.DeadlineDate,
		"subject_id":    newData.SubjectID,
		"subject_name":  newData.SubjectName,
	}

	err := s.repo.UpdateByID(id, data)
	return err
}

func (s *eventService) DeleteByID(id int) error {
	err := s.repo.DeleteByID(id)
	return err
}
