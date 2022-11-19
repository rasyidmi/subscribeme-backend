package main

// main
// controller
// service
// repository
// db

import (
	"net/http"
	"projects-subscribeme-backend/config"
	"projects-subscribeme-backend/controller"
	"time"

	timeout "github.com/s-wijaya/gin-timeout"

	"github.com/gin-gonic/gin"
)

func main() {
	config.SetupDatabase()
	firebaseAuth := config.SetupFirebase()

	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})
	router.Use(timeout.TimeoutHandler(5*time.Second, http.StatusRequestTimeout, gin.H{"data": "Request Timeout"}))

	v1Api := router.Group("/v1")

	v1Api.POST("/pengguna/daftar", controller.CreateUser)
	v1Api.GET("/pengguna/:id", controller.FindUserByID)

	v1Api.GET("/mata-kuliah", controller.GetAllSubjects)
	v1Api.POST("/mata-kuliah", controller.CreateSubject)

	v1Api.GET("/mata-kuliah/:id", controller.GetSubjectByID)
	v1Api.DELETE("/mata-kuliah/:id", controller.DeleteSubjectByID)
	v1Api.POST("/mata-kuliah/:id", controller.UpdateSubjectByID)

	v1Api.GET("/kelas/:id", controller.GetClassByID)

	v1Api.POST("/agenda", controller.CreateEvent)

	v1Api.GET("/agenda/:id", controller.GetEventByID)
	v1Api.POST("/agenda/:id", controller.UpdateEventByID)
	v1Api.DELETE("/agenda/:id", controller.DeleteEventByID)

	router.Run()
}
