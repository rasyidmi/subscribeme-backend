package routers

import (
	"projects-subscribeme-backend/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Event(g *gin.RouterGroup) {
	g.GET("/today-deadline", controllers.GetTodayDeadline)
	g.GET("/seven-day-deadline", controllers.GetSevenDayDeadline)
	g.GET("/:id", controllers.GetEventByID)
	g.POST("/create", controllers.CreateEvent)
	g.PUT("/update/:id", controllers.UpdateEventByID)
	g.DELETE("/delete/:id", controllers.DeleteEventByID)
}
