package initializers

import (
	"fmt"
	"log"
	"projects-subscribeme-backend/config"
	"projects-subscribeme-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() *gorm.DB {
	postgresConfig := config.LoadPostgresConfig()

	log.Println(postgresConfig)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		postgresConfig.Host,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Name,
		postgresConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	tx := db.Session(&gorm.Session{PrepareStmt: true})

	if err != nil {
		log.Fatal("Koneksi DB Gagal")
	}

	migrateDatabase(tx)

	return tx

}

func migrateDatabase(db *gorm.DB) {

	errMigrate := db.AutoMigrate(&models.User{})

	if errMigrate != nil {
		log.Fatal("Gagal Migrate")
	}

	log.Println("Migrate Berhasil!")

}
