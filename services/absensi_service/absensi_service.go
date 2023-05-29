package absensi_service

import (
	"projects-subscribeme-backend/dto/response"
)

type AbsensiService interface {
	GetClassScheduleByNpmMahasiswa(npm string) (*[]response.ClassScheduleResponse, error)
	GetClassScheduleDetailByScheduleId(scheduleId string) (*response.ClassScheduleResponse, error)
	GetClassScheduleByYearAndTerm(year, term string) (*[]response.ClassScheduleResponse, error)
	GetClassParticipantByClassCode(classCode string) (*[]response.ClassDetailResponse, error)
}
