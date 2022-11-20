package routers

import (
	"projects-subscribeme-backend/controller"

	"github.com/gin-gonic/gin"
)

func Class(g *gin.RouterGroup) {
	g.GET("/:id", controller.GetClassByID)
}
