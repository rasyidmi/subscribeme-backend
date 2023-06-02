package response

import (
	"projects-subscribeme-backend/models"

	"github.com/jinzhu/copier"
)

type CourseSceleResponse struct {
	ID              string `json:"id"`
	CourseSceleID   int64  `json:"course_secele_id"`
	CourseSceleName string `json:"course_scele_name"`
}

func NewCourseSceleResponse(model models.CourseScele) *CourseSceleResponse {
	var response CourseSceleResponse

	copier.Copy(&response, model)

	return &response
}
