package initializers

import (
	"log"
	"projects-subscribeme-backend/controllers/absensi_controller"
	"projects-subscribeme-backend/services/absensi_service"
	"time"

	"github.com/joho/godotenv"
)

//Absensi
var AbsensiController absensi_controller.AbsensiController
var absensiService absensi_service.AbsensiService

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println(err.Error())
	}
	time.Local = location
	initServices()
	initController()
}

func initServices() {
	absensiService = absensi_service.NewAbsensiService()
}

func initController() {
	AbsensiController = absensi_controller.NewAbsensiController(absensiService)
}
