package absensi_controller

import (
	"github.com/gin-gonic/gin"
)

type AbsensiController interface {
	GetClassScheduleByNpmMahasiswa(ctx *gin.Context)
	GetClassScheduleDetailByScheduleId(ctx *gin.Context)
	GetClassScheduleByYearAndTerm(ctx *gin.Context)
	GetClassParticipantByClassCode(ctx *gin.Context)
}
