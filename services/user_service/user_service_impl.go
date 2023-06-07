package user_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/repositories/user_repository"
	"time"

	"gorm.io/gorm"
)

type userService struct {
	repository user_repository.UserRepository
}

func NewUserService(repository user_repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(claims *helper.JWTClaim, payload payload.FcmPayload) (*response.LoginResponse, error) {
	user := models.User{
		Username: claims.Username,
		Npm:      claims.Npm,
		Role:     constant.UserRoleMahasiswa,
		FcmToken: payload.FcmToken,
	}

	user, err := s.repository.CreateUser(user)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	expirationTime := time.Now().Add(35064 * time.Hour)

	token, err := helper.RefreshJWT(claims, user.Role.String(), expirationTime)
	if err != nil {

		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	go func() {
		var data = map[string]interface{}{}
		data["npm"] = claims.Npm
		data["tahun"] = "2019"
		data["term"] = "2"

		models, err := helper.GetSiakngData[[]models.ClassSchedule](constant.GetClassDetailByNpmMahasiswa, data)
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
		}

		for _, v := range *models {
			layout := "Monday 15:04:05"
			day := helper.ConvertToWeekday(v.Day)
			str := fmt.Sprintf("%s %s", day, v.StartTime)

			t, err := time.Parse(layout, str)
			if err != nil {
				fmt.Println("Terjadi kesalahan dalam parsing waktu:", err)
				return
			}

			t = t.Add(-15 * time.Minute)

			cronExpression := fmt.Sprintf("%d %d * * %d", t.Minute(), t.Hour(), int(t.Weekday()))

			jsonBytes, err := json.Marshal(v)
			if err != nil {
				log.Println(string("\033[31m"), err.Error())
				return
			}

			helper.SchedulerEvent.ScheduleCron("ReminderClassWillStarted", string(jsonBytes), cronExpression, user.ID.String(), "")
		}

	}()

	return response.NewLoginResponse(token, nil), nil

}

func (s *userService) LoginFromSSOUI(ticket string) (*response.LoginResponse, error) {

	user, err := helper.ValidateSSOTicket(ticket, "http://localhost:8080/api/v1/login/sso")
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	expirationTime := time.Now().Add(2 * time.Hour)

	token, err := helper.GenerateJWT(*user, "Mahasiswa", expirationTime)
	if err != nil {

		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	return response.NewLoginResponse(token, nil), nil
}

func (s *userService) Login(payload payload.SSOPayload) (*response.LoginResponse, error) {
	sso, err := helper.ValidateSSOTicket(payload.Ticket, payload.ServiceUrl)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	role := "Mahasiswa"
	isExists := true

	user, err := s.repository.GetUserByUsername(sso.AuthenticationSuccess.User)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isExists = false
		} else {
			log.Println(string("\033[31m"), err.Error())
			return nil, errors.New("500")
		}

	}

	var expirationTime time.Time

	if isExists {
		role = user.Role.String()
		expirationTime = time.Now().Add(35064 * time.Hour)
	} else {
		expirationTime = time.Now().Add(2 * time.Hour)
	}

	token, err := helper.GenerateJWT(*sso, role, expirationTime)
	if err != nil {

		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	return response.NewLoginResponse(token, &isExists), nil

}

func (s *userService) UpdateFcmTokenUser(claims *helper.JWTClaim, payload payload.FcmPayload) (*response.UserResponse, error) {
	user, err := s.repository.UpdateFcmTokenUser(claims.Username, payload.FcmToken)
	if err != nil {

		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	return response.NewUserResponse(user), nil
}
