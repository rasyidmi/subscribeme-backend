package absensi_service

import (
	"encoding/json"
	"errors"
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	absensi_repository "projects-subscribeme-backend/repositories/absence_repository"
	"projects-subscribeme-backend/repositories/user_repository"
	"time"

	"github.com/google/uuid"
	"github.com/jftuga/geodist"
	"gorm.io/gorm"
)

type absensiService struct {
	repository     absensi_repository.AbsensiRepository
	userRepository user_repository.UserRepository
}

func NewAbsensiService(repository absensi_repository.AbsensiRepository, userRepository user_repository.UserRepository) AbsensiService {
	return &absensiService{repository: repository, userRepository: userRepository}
}

func (s *absensiService) CreateAbsenceSession(payload payload.ClassAbsenceSessionPayload, claims *helper.JWTClaim) (*response.ClassAbsenceSessionResponse, error) {
	_, err := s.CheckAbsenceIsOpen(payload.ClassCode)

	classSessionCheck := false
	if err != nil {
		if err.Error() == "404" {
			classSessionCheck = true
		} else {
			log.Println(string("\033[31m"), err.Error())
			return nil, err
		}
	}

	if !classSessionCheck {
		log.Println(string("\033[31m"), errors.New("403"))
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

		model.Latitude = &payload.Latitude
		model.Longitude = &payload.Longitude
		model.GeoRadius = &payload.GeoRadius
		model.IsGeofence = payload.IsGeofence
	}

	absenceSession, err := s.repository.CreateAbsenceSession(model)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	go s.createAbsence(absenceSession, payload.StartTime)

	return response.NewClassAbsenceSessionResponse(absenceSession, false, false), nil

}

func (s *absensiService) UpdateAbsenceSession(payload payload.ClassAbsenceSessionPayload, claims *helper.JWTClaim, id string) (*response.ClassAbsenceSessionResponse, error) {
	model := models.ClassAbsenceSession{
		TeacherName: claims.Nama,
		ClassCode:   payload.ClassCode,
		StartTime:   payload.StartTime,
		EndTime:     payload.StartTime.Add(time.Minute * time.Duration(payload.Duration)),
		Latitude:    nil,
		Longitude:   nil,
		GeoRadius:   nil,
		IsGeofence:  false,
	}

	if payload.IsGeofence {
		if payload.Latitude == 0 || payload.Longitude == 0 || payload.GeoRadius == 0 {
			return nil, errors.New("400")
		}

		model.Latitude = &payload.Latitude
		model.Longitude = &payload.Longitude
		model.GeoRadius = &payload.GeoRadius
		model.IsGeofence = payload.IsGeofence
	}

	absenceSession, err := s.repository.UpdateAbsenceSession(model, id)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	go s.createAbsence(absenceSession, payload.StartTime)

	return response.NewClassAbsenceSessionResponse(absenceSession, false, false), nil
}

func (s *absensiService) createAbsence(payload models.ClassAbsenceSession, absenceOpenTime time.Time) (bool, error) {
	model, err := s.GetClassParticipantByClassCode(payload.ClassCode)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return false, err
	}

	classParticipant := *model

	absences := []models.Absence{}
	for _, val := range classParticipant {
		absence := &models.Absence{
			ClassAbsenceSessionID: payload.ID.String(),
			StudentName:           val.Student[0].Name,
			StudentNpm:            val.Student[0].Npm,
			ClassDate:             absenceOpenTime,
			Present:               false,
		}

		user, err := s.userRepository.GetUserByNpm(val.Student[0].Npm)
		if err == nil {
			loc, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				log.Println(string("\033[31m"), err.Error())
				return false, err
			}

			payload.StartTime = payload.StartTime.In(loc)
			payload.EndTime = payload.EndTime.In(loc)

			jsonBytes, err := json.Marshal(payload)
			if err != nil {
				log.Println(string("\033[31m"), err.Error())
				return false, err
			}

			endTime := payload.EndTime.Add(-time.Minute * 5)
			helper.SchedulerEvent.Schedule("ReminderAbsenceCanBeDone", string(jsonBytes), payload.StartTime, user.ID.String(), "")
			helper.SchedulerEvent.Schedule("ReminderAbsenceWillOver", string(jsonBytes), endTime, user.ID.String(), "")
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

	if absenceSession.EndTime.Before(time.Now()) && absenceSession.StartTime.After(time.Now()) {
		return nil, errors.New("absence slot not open yet or absence slot has ended")
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
		//TODO Check Virsanti
		classRoomLoc := geodist.Coord{Lat: *absenceSession.Latitude, Lon: *absenceSession.Longitude}
		studentLoca := geodist.Coord{Lat: payload.Latitude, Lon: payload.Longitude}

		_, km, err := geodist.VincentyDistance(classRoomLoc, studentLoca)
		if err != nil {
			return nil, errors.New("error to compute distance")
		}

		if km*1000 > *absenceSession.GeoRadius {
			return nil, errors.New("Your distance is too far from classroom")
		}

		absence.Latitude = payload.Latitude
		absence.Longitude = payload.Longitude

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

	return response.NewClassAbsenceSessionResponse(absenceClass, false, false), nil

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

func (s *absensiService) GetAbsenceSessionDetailByAbsenceSessionId(absenceSessionId string) (*[]response.ClassAbsenceSessionResponse, error) {

	absenceSessions, err := s.repository.GetAbsenceSessionByAbsenceSessionId(absenceSessionId)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewClassAbsenceSessionResponses(absenceSessions, true, true), nil

}

func (s *absensiService) GetAbsenceSessionByClassCode(classCode string) (*[]response.ClassAbsenceSessionResponse, error) {
	absenceSession, err := s.repository.GetAbsenceSessionByClassCode(classCode)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewClassAbsenceSessionResponses(absenceSession, true, false), nil
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

	models, err := helper.GetSiakngData[[]models.ClassSchedule](constant.GetClassDetailByNpmMahasiswa, data)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewClassDetailResponses(*models), nil
}

func (s *absensiService) GetClassDetailByNimDosen(nim string) (*[]response.CourseResponseDosen, error) {
	var data = map[string]interface{}{}

	data["nim"] = nim
	data["tahun"] = "2019"
	data["term"] = "2"

	models, err := helper.GetSiakngData[[]models.ClassSchedule](constant.GetClassDetailByNimDosen, data)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewCourceResponseDosen(*models), nil

}

func (s *absensiService) GetClassParticipantByClassCode(classCode string) (*[]response.ListStudentResponse, error) {
	var data = map[string]interface{}{}

	data["kd_kls"] = classCode

	models, err := helper.GetSiakngData[[]models.ClassDetail](constant.GetClassParticipantByClassCode, data)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewClassParticipantResponses(*models), nil
}
