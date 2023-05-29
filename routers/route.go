package routers

import (
	"net/http"
	"projects-subscribeme-backend/initializers"

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

	siakng := api.Group("/siakng")
	siakng.GET("/class/npm", initializers.AbsensiController.GetClassScheduleByNpmMahasiswa)
	siakng.GET("/class/:class_code", initializers.AbsensiController.GetClassScheduleDetailByScheduleId)
	siakng.GET("/class/schedule/:year/:term", initializers.AbsensiController.GetClassScheduleByYearAndTerm)
	siakng.GET("/class/participants/:class_code", initializers.AbsensiController.GetClassParticipantByClassCode)

	return router

}
