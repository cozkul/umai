package models

import (
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
