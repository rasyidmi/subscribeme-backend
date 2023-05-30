package absensi_service

import (
	"errors"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/helper"
)

type absensiService struct {
}

func NewAbsensiService() AbsensiService {
	return &absensiService{}
}

func (s *absensiService) GetClassScheduleByNpmMahasiswa(npm string) (*[]response.ClassScheduleResponse, error) {
	var data = map[string]interface{}{}

	data["npm"] = npm
	data["tahun"] = "2019"

	// month := time.Now().Month()
	// if month < time.July {
	// 	data["term"] = 1
	// } else {
	// 	data["term"] = 2
	// }

	data["term"] = "2"

	models, err := helper.GetSiakngData[[]models.ClassSchedule](constant.GetClassScheduleByNpmMahasiswa, data)

	if err != nil {
		return nil, err
	}

	if len(*models) == 0 {
		return nil, errors.New("404")
	}

	return response.NewClassScheduleResponses(*models), nil
}

func (s *absensiService) GetClassScheduleDetailByScheduleId(scheduleId string) (*response.ClassScheduleResponse, error) {
	var data = map[string]interface{}{}

	data["schedule_id"] = scheduleId

	models, err := helper.GetSiakngData[models.ClassSchedule](constant.GetClassScheduleByScheduleId, data)

	if err != nil {
		return nil, err
	}

	if models.ScheduleUrl == "" {
		return nil, errors.New("404")
	}

	return response.NewClassScheduleResponse(*models), nil

}

func (s *absensiService) GetClassParticipantByClassCode(classCode string) (*[]response.ClassDetailResponse, error) {
	var data = map[string]interface{}{}

	data["kd_kls"] = classCode

	models, err := helper.GetSiakngData[[]models.ClassDetail](constant.GetClassParticipantByClassCode, data)

	if err != nil {
		return nil, err
	}

	if len(*models) == 0 {
		return nil, errors.New("404")
	}

	return response.NewClassDetailResponses(*models), nil
}

func (s *absensiService) GetClassScheduleByYearAndTerm(year, term string) (*[]response.ClassScheduleResponse, error) {
	var data = map[string]interface{}{}

	data["year"] = year
	data["term"] = term

	models, err := helper.GetSiakngData[[]models.ClassSchedule](constant.GetClassScheduleByYearAndTerm, data)

	if err != nil {
		return nil, err
	}

	if len(*models) == 0 {
		return nil, errors.New("404")
	}

	return response.NewClassScheduleResponses(*models), nil

}
