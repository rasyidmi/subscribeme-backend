package routers

import (
	"projects-subscribeme-backend/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Class(g *gin.RouterGroup) {
	g.GET("/:id", controllers.GetClassByID)
}
