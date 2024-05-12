package database

import (
	"fmt"
	"log"
	"os"

	"github.com/cozkul/umai/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDb() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Canada/Pacific",
		config.Config.DBHost,
		config.Config.DBUserName,
		config.Config.DBUserPassword,
		config.Config.DBName,
		config.Config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(2)
	}

	log.Println("Database connected.")
	DB = db
	db.Logger = logger.Default.LogMode(logger.Info)
}
