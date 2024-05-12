package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username,omitempty" gorm:"text;not null;default:null"`
	Email    string `json:"email,omitempty" gorm:"text;not null;default:null"`
	Password string `json:"password,omitempty" gorm:"text;not null;default:null"`
	Verified *bool  `gorm:"not null;default:false"`
	Systems  []System
	Planets  []Planet
}

type FilteredUser struct {
	ID       uint
	UserName string `json:"username"`
	Email    string `json:"email"`
}
