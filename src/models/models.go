package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model

	Alias  string `gorm:"not null" json:"alias"`
	Url    string `gorm:"not null" json:"url"`
	Visits uint   `gorm:"not null;default:0" json:"visits"`
}
