package absensi_controller

import (
	"github.com/gin-gonic/gin"
)

type AbsensiController interface {
	CreateAbsenceSession(ctx *gin.Context)
	UpdateAbsenceSession(ctx *gin.Context)
	UpdateAbsence(ctx *gin.Context)
	CheckAbsenceIsOpen(ctx *gin.Context)
	GetAbsenceByClassCodeAndNpm(ctx *gin.Context)
	GetClassDetailByNimDosen(ctx *gin.Context)
	GetAbsenceSessionDetailByAbsenceSessionId(ctx *gin.Context)
	GetAbsenceSessionByClassCode(ctx *gin.Context)
	GetClassDetailByNpmMahasiswa(ctx *gin.Context)
	GetClassParticipantByClassCode(ctx *gin.Context)
}
