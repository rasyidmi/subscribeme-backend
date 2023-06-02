package course_service

import (
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
)

type CourseService interface {
	GetCoursesByUsername(claims *helper.JWTClaim) (*[]response.CourseMoodleResponse, error)
}
