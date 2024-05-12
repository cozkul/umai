package models

type SystemType string

const (
	PlanetarySystemType SystemType = "planetary system"
	NebulaSystemType    SystemType = "nebula"
	BlackHoleSystemType SystemType = "black hole"
)

type System struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	UserID     *uint
	SystemType SystemType
	Stars      []Star      `gorm:"foreignKey:SystemID"`
	Planets    []Planet    `gorm:"foreignKey:SystemID"`
	Nebula     Nebula      `gorm:"foreignKey:SystemID"`
	BlackHoles []BlackHole `gorm:"foreignKey:SystemID"`
}

type HyperLane struct {
	ID                uint   `gorm:"primaryKey"`
	SourceNodeID      uint   `gorm:"not null"`
	SourceNode        System `gorm:"foreignKey:SourceNodeID"`
	DestinationNodeID uint   `gorm:"not null"`
	DestinationNode   System `gorm:"foreignKey:DestinationNodeID"`
	EdgeWeight        int    `gorm:"not null"`
}
