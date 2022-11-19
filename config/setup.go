package config

import (
	"fmt"
	"log"
	"os"
	"projects-subscribeme-backend/controller"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/repository"
	"projects-subscribeme-backend/service"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var subjectServ service.SubjectService
var classServ service.ClassService
var eventServ service.EventService
var userServ service.UserService

var subjectRepo repository.SubjectRepository
var classRepo repository.ClassRepository
var eventRepo repository.EventRepository
var userRepo repository.UserRepository

func SetupDatabase() {
	connectDatabase()
	createRepositories()
	createServices()
	createControllers()
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func connectDatabase() {
	var dbHost = getEnvVariable("DB_HOST")
	var dbName = getEnvVariable("DB_NAME")
	var dbUser = getEnvVariable("DB_USER")
	var dbPass = getEnvVariable("DB_PASS")

	dsn := fmt.Sprint("host=", dbHost, " user=", dbUser, " password=", dbPass, " dbname=", dbName, " port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("!!!!!!!!!!!!!!!!!")
		log.Fatal("error on db setup")
		log.Fatal("!!!!!!!!!!!!!!!!!")
	}
	db.AutoMigrate(&models.Subject{}, &models.Class{}, &models.Event{}, &models.User{})

	DB = db
}

func createRepositories() {
	subjectRepo = repository.CreateSubjectRepository(DB)
	classRepo = repository.CreateClassRepository(DB)
	eventRepo = repository.CreateEventRepository(DB)
	userRepo = repository.CreateUserRepository(DB)
}

func createServices() {
	subjectServ = service.CreateSubjectService(subjectRepo)
	classServ = service.CreateClassService(classRepo)
	eventServ = service.CreateEventService(eventRepo)
	userServ = service.CreateUserService(userRepo)
}

func createControllers() {
	controller.CreateSubjectController(subjectServ)
	controller.CreateClassController(classServ)
	controller.CreateEventController(eventServ)
	controller.CreateUserController(userServ)
}
