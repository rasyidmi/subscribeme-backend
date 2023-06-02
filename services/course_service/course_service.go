package course_service

import (
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
)

type CourseService interface {
	GetCoursesByUsername(claims *helper.JWTClaim) (*[]response.CourseMoodleResponse, error)
	SubscribeCourse(claims *helper.JWTClaim, payload payload.ChooseCourse) (*response.CourseSceleResponse, error)
}
