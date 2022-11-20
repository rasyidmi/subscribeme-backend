package routers

import (
	"projects-subscribeme-backend/controller"

	"github.com/gin-gonic/gin"
)

func User(g *gin.RouterGroup) {
	g.POST("/register", controller.CreateUser)
	g.GET("/:id", controller.FindUserByID)
}
