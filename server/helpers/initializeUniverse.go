package helpers

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
)

func InitializeUniverse() error {
	var status models.DatabaseStatus
	if err := database.DB.First(&status).Error; err != nil {
		status = models.DatabaseStatus{
			Generated: false,
		}
		database.DB.Create(&status)
	}

	if !status.Generated {
		if err := generateUniverse(20, 2, 0.2); err != nil {
			return err
		}
		if err := database.DB.Model(&status).Update("Generated", true).Error; err != nil {
			return err
		}
	}
	return nil
}

// Function to generate a random system type
func randomSystemType() models.SystemType {
	randomNum := rand.Intn(1000) + 1
	if randomNum <= 900 {
		return "planetary system"
	} else if randomNum <= 985 {
		return "nebula"
	} else {
		return "black hole"
	}
}

// Function to generate a random star count
func randomStarCount() int {
	randomNum := rand.Intn(100) + 1
	if randomNum <= 70 {
		return 1
	} else if randomNum <= 95 {
		return 2
	} else {
		return 3
	}
}

// Function to generate a random planet count
func randomPlanetCount() int {
	return rand.Intn(6) + 5
}

// Function to generate a random black hole count
func randomBlackHoleCount() int {
	randomNum := rand.Intn(100) + 1

	if randomNum <= 80 {
		return 1
	} else if randomNum <= 98 {
		return 2
	} else {
		return 3
	}
}

// Function to generate a random moon count
func randomMoonCount() int {
	randomNum := rand.Intn(20) + 1
	if randomNum <= 18 {
		return 0
	} else if randomNum <= 19 {
		return 1
	} else {
		return 2
	}
}

// Function to generate a planetary system
func generatePlanetarySystem(system *models.System) {
	for i := 1; i <= randomStarCount(); i++ {
		star := models.Star{
			Celestial: models.Celestial{
				Name: fmt.Sprintf("Star %d", i),
			},
			Luminosity:  rand.Intn(38) + 19,
			Temperature: rand.Intn(1500) + 1500,
		}
		system.Stars = append(system.Stars, star)
	}

	n := randomPlanetCount()
	for i := 1; i <= n; i++ {
		planet := models.Planet{
			Celestial: models.Celestial{
				Name: fmt.Sprintf("Planet %d", i),
			},
			SurfaceTemperature: int(float32(i)/float32(n)*rand.Float32()*200) + 0,
			Size:               rand.Intn(10) + 5,
		}
		system.Planets = append(system.Planets, planet)

		for j := 1; j <= randomMoonCount(); j++ {
			moon := models.Moon{
				Celestial: models.Celestial{
					Name: fmt.Sprintf("Moon %d", i),
				},
				Size: rand.Intn(10) + 5,
			}
			planet.Moons = append(planet.Moons, moon)
		}
	}
}

// Function to generate a black hole system
func generateBlackHoleSystem(system *models.System) {
	blackHoleCount := randomBlackHoleCount()
	for i := 1; i <= blackHoleCount; i++ {
		blackHole := models.BlackHole{
			Celestial: models.Celestial{
				Name: fmt.Sprintf("Moon %d", i),
			},
			Size: rand.Intn(100) + 50,
		}
		system.BlackHoles = append(system.BlackHoles, blackHole)
	}
}

// Function to generate a nebula system
func generateNebulaSystem(system *models.System) {
	system.Nebula = models.Nebula{
		Celestial: models.Celestial{
			Name: "Nebula",
		},
	}
}

// Function to generate a Watts-Strogatz graph representing a universe
func generateUniverse(n int, k int, p float64) error {
	database.DB.Exec("DELETE FROM hyper_lanes")
	database.DB.Exec("DELETE FROM systems")

	for i := 1; i <= n; i++ {
		systemType := randomSystemType()

		switch systemType {
		case "planetary system":
			system := models.System{
				Name:       fmt.Sprintf("Planetary System %d", i),
				SystemType: models.PlanetarySystemType,
			}
			generatePlanetarySystem(&system)
			log.Println("adding planetary system")
			database.DB.Create(&system)
			log.Println("adding planetary system complete")
		case "nebula":
			system := models.System{
				Name:       fmt.Sprintf("Nebula System %d", i),
				SystemType: models.NebulaSystemType,
			}
			generateNebulaSystem(&system)
			database.DB.Create(&system)
		case "black hole":
			system := models.System{
				Name:       fmt.Sprintf("Black Hole System %d", i),
				SystemType: models.BlackHoleSystemType,
			}
			generateBlackHoleSystem(&system)
			database.DB.Create(&system)
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			destinationNodeID := ((i + j - 1) % n) + 1
			edgeWeight := rand.Intn(5) + 10
			var sourceNode, destinationNode models.System
			database.DB.First(&sourceNode, i)
			database.DB.First(&destinationNode, destinationNodeID)
			database.DB.Create(&models.HyperLane{
				SourceNode:      sourceNode,
				DestinationNode: destinationNode,
				EdgeWeight:      edgeWeight},
			)
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if rand.Float64() < p {
				destinationNodeID := uint(rand.Intn(n) + 1)

				// Check if the destination node is not equal to the source node
				// and if an edge doesn't already exist between them
				var count int64
				database.DB.Model(&models.HyperLane{}).
					Where("source_node_id = ? AND destination_node_id = ?", i, destinationNodeID).
					Count(&count)
				if count == 0 && i != int(destinationNodeID) {
					database.DB.Model(&models.HyperLane{}).
						Where("source_node_id = ?", i).
						Update("destination_node_id", destinationNodeID)
				}
			}
		}
	}
	return nil
}
