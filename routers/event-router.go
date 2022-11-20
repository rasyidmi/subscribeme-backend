package routers

import (
	"projects-subscribeme-backend/controller"

	"github.com/gin-gonic/gin"
)

func Event(g *gin.RouterGroup) {
	g.GET("/:id", controller.GetEventByID)
	g.POST("/create", controller.CreateEvent)
	g.PUT("/update/:id", controller.UpdateEventByID)
	g.DELETE("/delete/:id", controller.DeleteEventByID)
}
