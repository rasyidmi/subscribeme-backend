package routers

import (
	"projects-subscribeme-backend/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func User(g *gin.RouterGroup) {
	g.POST("/login", controllers.Login)
	g.POST("/autoLogin", controllers.AutoLogin)
	g.POST("/register", controllers.CreateUser)
	g.POST("/refresh", controllers.RefreshToken)
	g.GET("/:id", controllers.FindUserByID)
}
