package absensi_service

import (
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	"time"
)

type AbsensiService interface {
	//Mahasiswa
	createAbsence(payload models.ClassAbsenceSession, absenceOpenTime time.Time) (bool, error)
	UpdateAbsence(payload payload.AbsencePayload, claims *helper.JWTClaim) (*response.AbsenceResponse, error)
	CheckAbsenceIsOpen(classCode string) (*response.ClassAbsenceSessionResponse, error)
	GetAbsenceByClassCodeAndNpm(classCode string, claims *helper.JWTClaim) (*[]response.AbsenceResponse, error)
	GetClassDetailByNpmMahasiswa(npm string) (*[]response.ClassDetailResponse, error)

	//Dosen
	CreateAbsenceSession(payload payload.ClassAbsenceSessionPayload, claims *helper.JWTClaim) (*response.ClassAbsenceSessionResponse, error)
	UpdateAbsenceSession(payload payload.ClassAbsenceSessionPayload, claims *helper.JWTClaim, id string) (*response.ClassAbsenceSessionResponse, error)
	GetClassDetailByNimDosen(nim string) (*[]response.CourseResponseDosen, error)
	GetClassParticipantByClassCode(classCode string) (*[]response.ListStudentResponse, error)
	GetAbsenceSessionByClassCode(classCode string) (*[]response.ClassAbsenceSessionResponse, error)
	GetAbsenceSessionDetailByAbsenceSessionId(absenceSessionId string) (*[]response.ClassAbsenceSessionResponse, error)
}
