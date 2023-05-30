package absensi_controller

import (
	"errors"
	"net/http"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/services/absensi_service"

	"github.com/gin-gonic/gin"
)

type absensiController struct {
	service absensi_service.AbsensiService
}

func NewAbsensiController(service absensi_service.AbsensiService) AbsensiController {
	return &absensiController{service: service}
}

func (c *absensiController) GetClassScheduleByNpmMahasiswa(ctx *gin.Context) {
	claims := helper.GetTokenClaims(ctx)
	data, err := c.service.GetClassScheduleByNpmMahasiswa(claims.Npm)
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

func (c *absensiController) GetClassScheduleDetailByScheduleId(ctx *gin.Context) {
	classCode := ctx.Param("class_code")

	data, err := c.service.GetClassScheduleDetailByScheduleId(classCode)
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

func (c *absensiController) GetClassScheduleByYearAndTerm(ctx *gin.Context) {
	year := ctx.Param("year")
	term := ctx.Param("term")

	data, err := c.service.GetClassScheduleByYearAndTerm(year, term)
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

func (c *absensiController) GetClassParticipantByClassCode(ctx *gin.Context) {
	classCode := ctx.Param("class_code")

	data, err := c.service.GetClassParticipantByClassCode(classCode)
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
