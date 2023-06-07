package absensi_controller

import (
	"errors"
	"net/http"
	"projects-subscribeme-backend/dto/payload"
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

func (c *absensiController) CreateAbsenceSession(ctx *gin.Context) {
	var payload payload.ClassAbsenceSessionPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := helper.GetTokenClaims(ctx)

	data, err := c.service.CreateAbsenceSession(payload, claims)
	if err != nil {
		if err.Error() == "400" {
			response.Error(ctx, "failed", http.StatusBadRequest, errors.New("payload error"))
			ctx.Abort()
			return
		} else if err.Error() == "403" {
			response.Error(ctx, "failed", http.StatusForbidden, errors.New("there's still another absence session that are open"))
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)

}

func (c *absensiController) UpdateAbsenceSession(ctx *gin.Context) {
	var payload payload.ClassAbsenceSessionPayload

	absenceSessionId := ctx.Param("absence_session_id")

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := helper.GetTokenClaims(ctx)

	data, err := c.service.UpdateAbsenceSession(payload, claims, absenceSessionId)
	if err != nil {
		if err.Error() == "400" {
			response.Error(ctx, "failed", http.StatusBadRequest, errors.New("payload error"))
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)

}

func (c *absensiController) UpdateAbsence(ctx *gin.Context) {
	var payload payload.AbsencePayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := helper.GetTokenClaims(ctx)

	data, err := c.service.UpdateAbsence(payload, claims)
	if err != nil {
		if err.Error() == "400" {
			response.Error(ctx, "failed", http.StatusBadRequest, errors.New("payload error"))
			ctx.Abort()
			return
		} else if err.Error() == "absence slot not open yet or absence slot has ended" {
			response.Error(ctx, "failed", http.StatusForbidden, err)
			ctx.Abort()
			return
		} else if err.Error() == "Your distance is too far from classroom" {
			response.Error(ctx, "failed", http.StatusForbidden, err)
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)

}

func (c *absensiController) CheckAbsenceIsOpen(ctx *gin.Context) {
	classCode := ctx.Param("class_code")

	data, err := c.service.CheckAbsenceIsOpen(classCode)
	if err != nil {
		if err.Error() == "404" {
			response.Error(ctx, "failed", http.StatusNotFound, errors.New("No absence session opened"))
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)
}

func (c *absensiController) GetAbsenceByClassCodeAndNpm(ctx *gin.Context) {
	classCode := ctx.Param("class_code")
	claims := helper.GetTokenClaims(ctx)

	data, err := c.service.GetAbsenceByClassCodeAndNpm(classCode, claims)
	if err != nil {
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)
}

func (c *absensiController) GetAbsenceSessionByClassCode(ctx *gin.Context) {
	classCode := ctx.Param("class_code")

	data, err := c.service.GetAbsenceSessionByClassCode(classCode)
	if err != nil {
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)
}

func (c *absensiController) GetAbsenceSessionDetailByAbsenceSessionId(ctx *gin.Context) {
	absenceSessionId := ctx.Param("absence_session_id")

	data, err := c.service.GetAbsenceSessionDetailByAbsenceSessionId(absenceSessionId)
	if err != nil {
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, data)

}

func (c *absensiController) GetClassDetailByNpmMahasiswa(ctx *gin.Context) {
	claims := helper.GetTokenClaims(ctx)
	data, err := c.service.GetClassDetailByNpmMahasiswa(claims.Npm)
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

func (c *absensiController) GetClassDetailByNimDosen(ctx *gin.Context) {
	claims := helper.GetTokenClaims(ctx)
	data, err := c.service.GetClassDetailByNimDosen(claims.Npm)
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
