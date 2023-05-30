package routers

import (
	"net/http"
	"projects-subscribeme-backend/initializers"
	"projects-subscribeme-backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPut},
		AllowHeaders:     []string{"Content-Type", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	api := router.Group("/api/v1")
	api.GET("/login/sso", initializers.UserController.LoginWithSSO)

	user := api.Group("/user")
	user.POST("", initializers.UserController.CreateUser)

	siakng := api.Group("/siakng")
	siakng.GET("/class/npm", middlewares.Auth("Mahasiswa"), initializers.AbsensiController.GetClassScheduleByNpmMahasiswa)
	siakng.GET("/class/:class_code", middlewares.Auth("Mahasiswa"), initializers.AbsensiController.GetClassScheduleDetailByScheduleId)
	siakng.GET("/class/schedule/:year/:term", middlewares.Auth("Mahasiswa"), initializers.AbsensiController.GetClassScheduleByYearAndTerm)
	siakng.GET("/class/participants/:class_code", middlewares.Auth("Mahasiswa"), initializers.AbsensiController.GetClassParticipantByClassCode)

	return router

}
