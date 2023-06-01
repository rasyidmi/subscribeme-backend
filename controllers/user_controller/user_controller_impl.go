package user_controller

import (
	"errors"
	"net/http"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/services/user_service"

	"github.com/gin-gonic/gin"
)

type userContrroler struct {
	service user_service.UserService
}

func NewUserController(service user_service.UserService) UserController {
	return &userContrroler{service: service}
}

func (c *userContrroler) CreateUser(ctx *gin.Context) {
	var payload payload.FcmPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := helper.GetTokenClaims(ctx)

	user, err := c.service.CreateUser(claims, payload)
	if err != nil {
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)
}

func (c *userContrroler) LoginWithSSO(ctx *gin.Context) {
	ticket := ctx.Query("ticket")

	token, err := c.service.LoginFromSSOUI(ticket)
	if err != nil {
		if err.Error() == "401" {
			response.Error(ctx, "failed", http.StatusUnauthorized, errors.New("You're not Fasilkom"))
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, token)
}

func (c *userContrroler) Login(ctx *gin.Context) {
	var payload payload.SSOPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	token, err := c.service.Login(payload)
	if err != nil {
		if err.Error() == "401" {
			response.Error(ctx, "failed", http.StatusUnauthorized, errors.New("You're not Fasilkom"))
			ctx.Abort()
			return
		} else if err.Error() == "500" {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
			ctx.Abort()
			return
		}
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, token)

}

func (c *userContrroler) UpdateFcmTokenUser(ctx *gin.Context) {
	var payload payload.FcmPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := helper.GetTokenClaims(ctx)

	user, err := c.service.UpdateFcmTokenUser(claims, payload)
	if err != nil {
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)
}
