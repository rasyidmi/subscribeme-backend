package config

import (
	"fmt"
	"log"
	"os"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/controllers"
	"projects-subscribeme-backend/pkg/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var subjectModel models.SubjectModel
var classModel models.ClassModel
var eventModel models.EventRepository
var userModel models.UserModel

func SetupDatabase() {
	connectDatabase()
	createRepositories()
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
	db.AutoMigrate(&modelsConfig.Subject{}, &modelsConfig.Class{}, &modelsConfig.Event{}, &modelsConfig.User{})

	DB = db
}

func createRepositories() {
	subjectModel = models.CreateSubjectRepository(DB)
	classModel = models.CreateClassRepository(DB)
	eventModel = models.CreateEventRepository(DB)
	userModel = models.CreateUserModel(DB)
}

func createControllers() {
	controllers.CreateSubjectController(subjectModel)
	controllers.CreateClassController(classModel)
	controllers.CreateEventController(eventModel)
	controllers.CreateUserController(userModel)
}
