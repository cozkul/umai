package helpers

import (
	"math/rand"

	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
)

func InitializeUser(newUser *models.User) error {
	var firstSystem models.System
	result := database.DB.Preload("Planets").Where("user_id IS NULL AND system_type = ?", "planetary system").Order("RANDOM()").Limit(1).Find(&firstSystem)
	if result.Error != nil {
		return result.Error
	}

	firstPlanet := firstSystem.Planets[rand.Intn(len(firstSystem.Planets))]

	firstPlanet.Buildings = append(firstPlanet.Buildings, models.PlanetaryBuilding{
		BuildingType: models.PlanetaryControlCenter,
		PlanetID:     firstPlanet.ID,
		Level:        1,
		ResourceTick: &models.TickEvent{},
	})

	firstPlanet.Population = 1

	firstPlanet.Resource = models.Resource{
		Metal:   1000,
		Crystal: 1000,
	}

	newUser.Planets = append(newUser.Planets, firstPlanet)
	newUser.Systems = append(newUser.Systems, firstSystem)

	return nil
}
