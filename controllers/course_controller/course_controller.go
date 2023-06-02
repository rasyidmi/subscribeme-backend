package course_controller

import "github.com/gin-gonic/gin"

type CourseController interface {
	GetCourseByUsername(ctx *gin.Context)
	SubscribeCourse(ctx *gin.Context)
	GetUserCourseByUsername(ctx *gin.Context)
	GetUserEventByCourseId(ctx *gin.Context)
	GetDeadlineTodayByUserId(ctx *gin.Context)
	GetDeadline7DaysAheadByUserId(ctx *gin.Context)
}
