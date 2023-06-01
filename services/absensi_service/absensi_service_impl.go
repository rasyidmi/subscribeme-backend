package absensi_service

import (
	"errors"
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	absensi_repository "projects-subscribeme-backend/repositories/absence_repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type absensiService struct {
	repository absensi_repository.AbsensiRepository
}

func NewAbsensiService(repository absensi_repository.AbsensiRepository) AbsensiService {
	return &absensiService{repository: repository}
}

func (s *absensiService) CreateAbsenceSession(payload payload.ClassAbsenceSessionPayload, claims *helper.JWTClaim) (*response.ClassAbsenceSessionResponse, error) {
	_, err := s.CheckAbsenceIsOpen(payload.ClassCode)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	model := models.ClassAbsenceSession{
		TeacherName: claims.Nama,
		ClassCode:   payload.ClassCode,
		StartTime:   payload.StartTime,
		EndTime:     payload.StartTime.Add(time.Minute * time.Duration(payload.Duration)),
	}

	if payload.IsGeofence {
		if payload.Latitude == 0 || payload.Longitude == 0 || payload.GeoRadius == 0 {
			return nil, errors.New("400")
		}

		model.Latitude = payload.Latitude
		model.Longitude = payload.Longitude
		model.GeoRadius = payload.GeoRadius
	}

	absenceSession, err := s.repository.CreateAbsenceSession(model)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	go s.createAbsence(absenceSession)

	return response.NewClassAbsenceSessionResponse(absenceSession), nil

}

func (s *absensiService) createAbsence(payload models.ClassAbsenceSession) (bool, error) {
	model, err := s.GetClassParticipantByClassCode(payload.ClassCode)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return false, err
	}

	classParticipant := *model

	absences := []models.Absence{}
	for _, val := range classParticipant[0].ListStudent {
		absence := &models.Absence{
			ClassAbsenceSessionID: payload.ID.String(),
			StudentName:           val.Student[0].Name,
			StudentNpm:            val.Student[0].Npm,
			Present:               false,
		}

		absences = append(absences, *absence)
	}

	_, err = s.repository.CreateAbsence(absences)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return false, err
	}

	return true, nil
}

func (s *absensiService) UpdateAbsence(payload payload.AbsencePayload, claims *helper.JWTClaim) (*response.AbsenceResponse, error) {

	absenceSession, err := s.repository.GetAbsenceSessionById(payload.ClassAbsenceSessionId)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	if absenceSession.EndTime.Before(time.Now()) {
		return nil, errors.New("403")
	}

	absence := models.Absence{
		DeviceCode:  payload.DeviceCode,
		Present:     true,
		PresentTime: time.Now(),
	}

	if absenceSession.IsGeofence {
		if payload.Latitude == 0 || payload.Longitude == 0 {
			return nil, errors.New("400")
		}
		absence.Latitude = payload.Latitude
		absence.Longitude = payload.Longitude

		//TODO Check Virsanti
	}

	update, err := s.repository.UpdateAbsence(absence, claims.Npm, payload.ClassAbsenceSessionId)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewAbsenceResponse(update), err

}

func (s *absensiService) CheckAbsenceIsOpen(classCode string) (*response.ClassAbsenceSessionResponse, error) {
	absenceClass, err := s.repository.GetIsOpenAbsenceSessionByClassCodeAndEndTime(classCode)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.ClassAbsenceSessionResponse{}, errors.New("404")
		}
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	if absenceClass.ID == uuid.Nil {
		return &response.ClassAbsenceSessionResponse{}, errors.New("404")
	}

	return response.NewClassAbsenceSessionResponse(absenceClass), nil

}

func (s *absensiService) GetAbsenceByClassCodeAndNpm(classCode string, claims *helper.JWTClaim) (*[]response.AbsenceResponse, error) {
	var absence []models.Absence

	if claims.Role == "Mahasiswa" {
		models, err := s.repository.GetAbsenceByClassCodeAndNpm(classCode, claims.Npm)
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
			return nil, err
		}
		absence = models
	} else {
		return nil, errors.New("Invalid role in class")
	}

	return response.NewAbsenceResponses(absence), nil
}

func (s *absensiService) GetAbsenceByAbsenceSessionId(absenceSessionId string) (*[]response.AbsenceResponse, error) {

	absences, err := s.repository.GetAbsenceByAbsenceSessionId(absenceSessionId)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewAbsenceResponses(absences), nil

}

func (s *absensiService) GetAbsenceSessionByClassCode(classCode string) (*[]response.ClassAbsenceSessionResponse, error) {
	absenceSession, err := s.repository.GetAbsenceSessionByClassCode(classCode)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewClassAbsenceSessionResponses(absenceSession), nil
}

func (s *absensiService) GetClassDetailByNpmMahasiswa(npm string) (*[]response.ClassDetailResponse, error) {
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
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	if len(*models) == 0 {
		return nil, errors.New("404")
	}

	return response.NewClassDetailResponses(*models), nil
}

// func (s *absensiService) GetClassScheduleDetailByScheduleId(scheduleId string) (*response.ClassScheduleResponse, error) {
// 	var data = map[string]interface{}{}

// 	data["schedule_id"] = scheduleId

// 	models, err := helper.GetSiakngData[models.ClassSchedule](constant.GetClassScheduleByScheduleId, data)

// 	if err != nil {
// 		log.Println(string("\033[31m"), err.Error())
// 		return nil, err
// 	}

// 	if models.ScheduleUrl == "" {
// 		log.Println(string("\033[31m"), err.Error())
// 		return nil, errors.New("404")
// 	}

// 	return response.NewClassScheduleResponse(*models), nil

// }

func (s *absensiService) GetClassParticipantByClassCode(classCode string) (*[]response.ClassDetailResponse, error) {
	var data = map[string]interface{}{}

	data["kd_kls"] = classCode

	models, err := helper.GetSiakngData[[]models.ClassDetail](constant.GetClassParticipantByClassCode, data)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	if len(*models) == 0 {
		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	return response.NewClassParticipantResponses(*models), nil
}

// func (s *absensiService) GetClassScheduleByYearAndTerm(year, term string) (*[]response.ClassScheduleResponse, error) {
// 	var data = map[string]interface{}{}

// 	data["year"] = year
// 	data["term"] = term

// 	models, err := helper.GetSiakngData[[]models.ClassSchedule](constant.GetClassScheduleByYearAndTerm, data)

// 	if err != nil {
// 		log.Println(string("\033[31m"), err.Error())
// 		return nil, err
// 	}

// 	if len(*models) == 0 {
// 		log.Println(string("\033[31m"), err.Error())
// 		return nil, errors.New("404")
// 	}

// 	return response.NewClassScheduleResponses(*models), nil

// }
