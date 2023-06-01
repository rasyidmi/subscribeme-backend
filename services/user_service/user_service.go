package user_service

import (
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
)

type UserService interface {
	LoginFromSSOUI(ticket string) (*response.LoginResponse, error)
	Login(payload payload.SSOPayload) (*response.LoginResponse, error)
	CreateUser(claims *helper.JWTClaim, payload payload.FcmPayload) (*response.LoginResponse, error)
	UpdateFcmTokenUser(claims *helper.JWTClaim, payload payload.FcmPayload) (*response.UserResponse, error)
}
