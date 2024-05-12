package models

type DatabaseStatus struct {
	ID        uint `gorm:"primaryKey"`
	Generated bool `gorm:"default:false"`
}
