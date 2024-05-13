package models

import (
	"math"

	"github.com/cozkul/umai/server/config"
	"gorm.io/gorm"
)

type PlanetaryBuildingType string
type LunarBuildingType string
type StarbaseBuildingType string

const (
	PlanetaryControlCenter    PlanetaryBuildingType = "planetary control center"
	PlanetaryFarm             PlanetaryBuildingType = "planetary farm"
	MetalMine                 PlanetaryBuildingType = "metal mine"
	CrystalMine               PlanetaryBuildingType = "crystal mine"
	QuantumEssenceSynthesizer PlanetaryBuildingType = "quantum essance synthesizer"
	FoodSilo                  PlanetaryBuildingType = "food silo"
	MetalStorage              PlanetaryBuildingType = "metal storage"
	CrystalReserve            PlanetaryBuildingType = "crystal reserve"
	QuantumEssenceTank        PlanetaryBuildingType = "quantum essance tank"
	SolarPlant                PlanetaryBuildingType = "solar plant"
	ResearchLab               PlanetaryBuildingType = "research lab"
	QuantumReactor            PlanetaryBuildingType = "quantum reactor"
)

const (
	LunarControlCenter LunarBuildingType = "lunar control center"
	LunarTelescope     LunarBuildingType = "lunar telescope"
)

const (
	StarbaseHull StarbaseBuildingType = "starbase hull"
	Shipyard     StarbaseBuildingType = "shipyard"
	TradeHub     StarbaseBuildingType = "trade hub"
)

type PlanetaryBuilding struct {
	gorm.Model
	BuildingType PlanetaryBuildingType `gorm:"not null"`
	PlanetID     uint                  `gorm:"not null"`
	Level        int                   `gorm:"not null"`
	ResourceTick *TickEvent            `gorm:"polymorphic:Parent"`
}

type LunarBuilding struct {
	gorm.Model
	BuildingType LunarBuildingType `gorm:"not null"`
	MoonID       uint              `gorm:"not null"`
	Level        int               `gorm:"not null"`
}

type StarbaseBuilding struct {
	gorm.Model
	BuildingType StarbaseBuildingType `gorm:"not null"`
	StarbaseID   uint                 `gorm:"not null"`
	Level        int                  `gorm:"not null"`
}

func (h *PlanetaryBuilding) GetProduction() Resource {
	switch h.BuildingType {
	case "planetary control center":
		return Resource{
			Food:           float32(h.Level*20) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			Metal:          float32(h.Level*40) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			Crystal:        float32(h.Level*20) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			QuantumEssence: 0,
			Energy:         10,
		}
	case "food silo":
		return Resource{
			Food:           float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			Metal:          0,
			Crystal:        0,
			QuantumEssence: 0,
			Energy:         0,
		}
	case "metal storage":
		return Resource{
			Food:           0,
			Metal:          float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			Crystal:        0,
			QuantumEssence: 0,
			Energy:         0,
		}
	case "crystal reserve":
		return Resource{
			Food:           0,
			Metal:          0,
			Crystal:        float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			QuantumEssence: 0,
			Energy:         0,
		}
	case "quantum essance tank":
		return Resource{
			Food:           0,
			Metal:          0,
			Crystal:        0,
			QuantumEssence: float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))) * config.HourlyToTick,
			Energy:         0,
		}
	default:
		return Resource{
			Food:           0,
			Metal:          0,
			Crystal:        0,
			QuantumEssence: 0,
			Energy:         0,
		}
	}
}

func (h *PlanetaryBuilding) GetCapacity() Resource {
	switch h.BuildingType {
	case "planetary control center":
		return Resource{
			Food:           float32(h.Level*1000 + 1000),
			Metal:          float32(h.Level*1000 + 1000),
			Crystal:        float32(h.Level*1000 + 1000),
			QuantumEssence: float32(h.Level*1000 + 1000),
			Energy:         0,
		}
	case "food silo":
		return Resource{
			Food:           float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))),
			Metal:          0,
			Crystal:        0,
			QuantumEssence: 0,
			Energy:         0,
		}
	case "metal storage":
		return Resource{
			Food:           0,
			Metal:          float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))),
			Crystal:        0,
			QuantumEssence: 0,
			Energy:         0,
		}
	case "crystal reserve":
		return Resource{
			Food:           0,
			Metal:          0,
			Crystal:        float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))),
			QuantumEssence: 0,
			Energy:         0,
		}
	case "quantum essance tank":
		return Resource{
			Food:           0,
			Metal:          0,
			Crystal:        0,
			QuantumEssence: float32(h.Level*1000) * float32(math.Pow(1.1, float64(h.Level))),
			Energy:         0,
		}
	default:
		return Resource{
			Food:           0,
			Metal:          0,
			Crystal:        0,
			QuantumEssence: 0,
			Energy:         0,
		}
	}
}
