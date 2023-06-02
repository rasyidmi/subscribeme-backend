package response

import (
	"projects-subscribeme-backend/models"

	"github.com/jinzhu/copier"
)

type CourseMoodleResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Assignments []CourseAssignment `json:"assignments,omitempty"`
}

type CourseAssignment struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	DueDate int64  `json:"duedate"`
}

type CourseQuizzez struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	TimeOpen int64  `json:"timeopen"`
}

func NewCourseMoodleResponse(model []models.CourseMoodle) *[]CourseMoodleResponse {
	var response []CourseMoodleResponse

	copier.Copy(&response, model)

	return &response
}
