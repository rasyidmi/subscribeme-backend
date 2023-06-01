package absensi_service

import (
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
)

type AbsensiService interface {
	CreateAbsenceSession(payload payload.ClassAbsenceSessionPayload, claims *helper.JWTClaim) (*response.ClassAbsenceSessionResponse, error)
	createAbsence(payload models.ClassAbsenceSession) (bool, error)
	UpdateAbsence(payload payload.AbsencePayload, claims *helper.JWTClaim) (*response.AbsenceResponse, error)
	CheckAbsenceIsOpen(classCode string) (*response.ClassAbsenceSessionResponse, error)
	GetAbsenceByClassCodeAndNpm(classCode string, claims *helper.JWTClaim) (*[]response.AbsenceResponse, error)

	GetAbsenceSessionByClassCode(classCode string) (*[]response.ClassAbsenceSessionResponse, error)
	GetAbsenceByAbsenceSessionId(absenceSessionId string) (*[]response.AbsenceResponse, error)

	GetClassDetailByNpmMahasiswa(npm string) (*[]response.ClassDetailResponse, error)
	GetClassParticipantByClassCode(classCode string) (*[]response.ClassDetailResponse, error)
	// GetClassScheduleDetailByScheduleId(scheduleId string) (*response.ClassScheduleResponse, error)
	// GetClassScheduleByYearAndTerm(year, term string) (*[]response.ClassScheduleResponse, error)

}
