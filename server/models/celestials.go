package models

type Celestial struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	SystemID uint   `gorm:"not null"`
}

type Planet struct {
	Celestial
	SurfaceTemperature int
	Size               int
	Population         int
	UserID             *uint
	Moons              []Moon
	Buildings          []PlanetaryBuilding
	Resource           Resource `gorm:"polymorphic:Parent"`
}

type Star struct {
	Celestial
	Luminosity  int
	Temperature int
}

type Moon struct {
	Celestial
	Size      int
	PlanetID  uint
	Buildings []LunarBuilding
	Resource  Resource `gorm:"polymorphic:Parent"`
}

type Starbase struct {
	Celestial
	Buildings []StarbaseBuilding
	Resource  Resource `gorm:"polymorphic:Parent"`
}

type Debris struct {
	Celestial
	ParentBody uint
	Resource   Resource `gorm:"polymorphic:Parent"`
}

type BlackHole struct {
	Celestial
	Size int
}

type Nebula struct {
	Celestial
}

func (p *Planet) GetMaxCapacity() Resource {
	var totalCapacity Resource
	for _, b := range p.Buildings {
		totalCapacity.AddValue(b.GetCapacity())
	}
	return totalCapacity
}
