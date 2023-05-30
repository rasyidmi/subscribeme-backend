package user_controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	LoginWithSSO(ctx *gin.Context)
	Login(ctx *gin.Context)
}
