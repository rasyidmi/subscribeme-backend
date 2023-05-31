package initializers

import (
	"log"
	"projects-subscribeme-backend/controllers/absensi_controller"
	"projects-subscribeme-backend/controllers/user_controller"
	absensi_repository "projects-subscribeme-backend/repositories/absence_repository"
	"projects-subscribeme-backend/repositories/user_repository"
	"projects-subscribeme-backend/services/absensi_service"
	"projects-subscribeme-backend/services/user_service"
	"time"

	"github.com/joho/godotenv"
)

//User
var UserController user_controller.UserController
var userService user_service.UserService
var userRepository user_repository.UserRepository

//Absensi
var AbsensiController absensi_controller.AbsensiController
var absensiService absensi_service.AbsensiService
var absensiRepository absensi_repository.AbsensiRepository

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
	initRepositories()
	initServices()
	initController()
}

func initRepositories() {
	db := connectDatabase()

	userRepository = user_repository.NewUserRepository(db)
	absensiRepository = absensi_repository.NewAbsenceRepository(db)
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
	absensiService = absensi_service.NewAbsensiService(absensiRepository)
}

func initController() {
	UserController = user_controller.NewUserController(userService)
	AbsensiController = absensi_controller.NewAbsensiController(absensiService)
}
