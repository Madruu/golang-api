package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	ID      int
	Name    string
	Number  string
	UserID  uint
	Balance float64
}
