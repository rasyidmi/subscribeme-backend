package response

import (
	"projects-subscribeme-backend/models"

	"github.com/jinzhu/copier"
)

type CourseSceleResponse struct {
	ID              string `json:"id"`
	CourseSceleID   int64  `json:"course_scele_id"`
	CourseSceleName string `json:"course_scele_name"`
}

func NewCourseSceleResponse(model models.CourseScele) *CourseSceleResponse {
	var response CourseSceleResponse

	copier.Copy(&response, model)

	return &response
}

func NewCourseSceleResponses(models []*models.CourseScele) *[]CourseSceleResponse {
	var responses []CourseSceleResponse

	copier.Copy(&responses, models)

	return &responses
}
