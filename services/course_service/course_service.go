package course_service

import (
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
)

type CourseService interface {
	GetCoursesByUsername(claims *helper.JWTClaim) (*[]response.CourseMoodleResponse, error)
	SubscribeCourse(claims *helper.JWTClaim, payload payload.ChooseCourse) (*response.CourseSceleResponse, error)
	UnsibscribeCourse(claims *helper.JWTClaim, payload payload.ChooseCourse) (*response.CourseSceleResponse, error)
	GetUserCourseByUsername(claims *helper.JWTClaim) (*[]response.CourseSceleResponse, error)
	GetUserEventByCourseId(claims *helper.JWTClaim, courseId string) (*[]response.UserEventResponse, error)
	GetDeadlineTodayByUserId(claims *helper.JWTClaim) (*[]response.UserEventResponse, error)
	GetDeadline7DaysAheadByUserId(claims *helper.JWTClaim) (*[]response.UserEventResponse, error)

	SetDeadlineReminder(claims *helper.JWTClaim, payload payload.ReminderPayload) (bool, error)
}
