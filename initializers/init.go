package initializers

import (
	"log"
	"projects-subscribeme-backend/controllers/absensi_controller"
	"projects-subscribeme-backend/controllers/course_controller"
	"projects-subscribeme-backend/controllers/user_controller"
	"projects-subscribeme-backend/helper"
	absensi_repository "projects-subscribeme-backend/repositories/absence_repository"
	"projects-subscribeme-backend/repositories/course_repository"
	"projects-subscribeme-backend/repositories/user_repository"
	"projects-subscribeme-backend/services/absensi_service"
	"projects-subscribeme-backend/services/course_service"
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

//Courses
var CourseController course_controller.CourseController
var courseService course_service.CourseService
var courseRepository course_repository.CourseRepository

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
	helper.InitFirebase()
	SetupScheduler()
	
}

func initRepositories() {
	db := connectDatabase()
	DB = db

	userRepository = user_repository.NewUserRepository(db)
	absensiRepository = absensi_repository.NewAbsenceRepository(db)
	courseRepository = course_repository.NewCourseRepository(db)
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
	absensiService = absensi_service.NewAbsensiService(absensiRepository)
	courseService = course_service.NewCourseService(courseRepository, userRepository)
}

func initController() {
	UserController = user_controller.NewUserController(userService)
	AbsensiController = absensi_controller.NewAbsensiController(absensiService)
	CourseController = course_controller.NewCourseController(courseService)
}
