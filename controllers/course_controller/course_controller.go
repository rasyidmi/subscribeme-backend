package course_controller

import "github.com/gin-gonic/gin"

type CourseController interface {
	GetCourseByUsername(ctx *gin.Context)
}
