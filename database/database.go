package database

import (
	"github.com/tutv/go-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!", err.Error())
		os.Exit(2)
		return
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	//Todo: add migrations

	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	Database = DbInstance{Db: db}
}
