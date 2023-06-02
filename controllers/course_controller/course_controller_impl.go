package course_controller

import (
	"errors"
	"net/http"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/services/course_service"

	"github.com/gin-gonic/gin"
)

type courseController struct {
	service course_service.CourseService
}

func NewCourseController(service course_service.CourseService) CourseController {
	return &courseController{service: service}
}

func (c *courseController) GetCourseByUsername(ctx *gin.Context) {

	claims := helper.GetTokenClaims(ctx)

	data, err := c.service.GetCoursesByUsername(claims)
	if err != nil {
		if err.Error() == "404" {
			response.Error(ctx, "failed", http.StatusNotFound, errors.New("Cant Fetch Data From SIAK-NG API"))
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)

}
