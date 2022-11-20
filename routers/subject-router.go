package routers

import (
	"projects-subscribeme-backend/controller"

	"github.com/gin-gonic/gin"
)

func Subject(g *gin.RouterGroup) {
	g.GET("/", controller.GetAllSubjects)
	g.GET("/:id", controller.GetSubjectByID)
	g.POST("/create", controller.CreateSubject)
	g.PUT("/update/:id", controller.UpdateSubjectByID)
	g.DELETE("/delete/:id", controller.DeleteSubjectByID)
}
