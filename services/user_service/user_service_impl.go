package user_service

import (
	"errors"
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/repositories/user_repository"

	"gorm.io/gorm"
)

type userService struct {
	repository user_repository.UserRepository
}

func NewUserService(repository user_repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(claims *helper.JWTClaim, payload payload.FcmPayload) (*response.UserResponse, error) {
	user := models.User{
		Username: claims.Username,
		Role:     constant.UserRoleMahasiswa,
		FcmToken: payload.FcmToken,
	}

	user, err := s.repository.CreateUser(user)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewUserResponse(user), nil

}

func (s *userService) LoginFromSSOUI(ticket string) (*response.LoginResponse, error) {

	user, err := helper.ValidateSSOTicket(ticket, "http://localhost:8080/api/v1/login/sso")
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	token, err := helper.GenerateJWT(*user, "Mahasiswa")
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
		}
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	if isExists {
		role = user.Role.String()
	}

	token, err := helper.GenerateJWT(*sso, role)
	if err != nil {

		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	return response.NewLoginResponse(token, &isExists), nil

}
