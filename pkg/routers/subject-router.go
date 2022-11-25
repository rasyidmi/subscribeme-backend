package routers

import (
	"projects-subscribeme-backend/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Subject(g *gin.RouterGroup) {
	g.GET("/", controllers.GetAllSubjects)
	g.GET("/:id", controllers.GetSubjectByID)
	g.POST("/create", controllers.CreateSubject)
	g.PUT("/update/:id", controllers.UpdateSubjectByID)
	g.DELETE("/delete/:id", controllers.DeleteSubjectByID)
}
