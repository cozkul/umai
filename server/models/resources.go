package models

type Resource struct {
	ID             uint `gorm:"primaryKey"`
	ParentID       uint
	ParentType     string
	Food           uint64
	Metal          uint64
	Crystal        uint64
	QuantumEssence uint64
}
