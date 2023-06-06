package response

import (
	"projects-subscribeme-backend/models"

	"github.com/jinzhu/copier"
)

type CourseMoodleResponse struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	IsSubscribed bool               `json:"is_subscribed"`
	Assignments  []CourseAssignment `json:"assignments,omitempty"`
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

func NewCourseMoodleResponses(model []models.CourseMoodle, subscribedCourses []*models.CourseScele) *[]CourseMoodleResponse {
	var response []CourseMoodleResponse

	copier.Copy(&response, model)

	for i := 0; i < len(response); i++ {
		for j := 0; j < len(subscribedCourses); j++ {
			if response[i].ID == subscribedCourses[j].CourseSceleID {
				response[i].IsSubscribed = true
			}
		}
	}

	return &response
}

func NewCourseMoodleResponse(model models.CourseMoodle) *CourseMoodleResponse {
	var response CourseMoodleResponse

	copier.Copy(&response, model)

	return &response
}
