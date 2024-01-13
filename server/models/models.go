package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"text;not null;default:null"`
	Email    string `json:"email" gorm:"text;not null;default:null"`
	Password string `json:"password" gorm:"text;not null;default:null"`
}
