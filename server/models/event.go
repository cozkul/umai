package models

import "gorm.io/gorm"

type TickEventType string

type TimedEvent struct {
	gorm.Model
	ParentID   uint
	ParentType string
}

type TickEvent struct {
	gorm.Model
	ParentID   uint
	ParentType string
}
