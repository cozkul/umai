package database

import (
	"log"

	"github.com/cozkul/umai/server/models"
)

func RunMigrations() error {
	// Runs queries
	log.Println("Running migrations.")

	DB.AutoMigrate(
		&models.DatabaseStatus{},
	)

	DB.AutoMigrate(
		&models.User{},
	)

	DB.AutoMigrate(
		&models.System{},
		&models.HyperLane{},
	)

	DB.AutoMigrate(
		&models.Resource{},
	)

	DB.AutoMigrate(
		&models.Planet{},
		&models.Star{},
		&models.Moon{},
		&models.Debris{},
		&models.BlackHole{},
		&models.Nebula{},
	)

	DB.AutoMigrate(
		&models.PlanetaryBuilding{},
		&models.LunarBuilding{},
		&models.StarbaseBuilding{},
	)

	return nil
}
