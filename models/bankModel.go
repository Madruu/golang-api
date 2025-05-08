package models

type Bank struct {
	//gorm.Model
	ID      uint `json:"id" gorm:"primaryKey"`
	Name    string
	Number  string
	UserID  uint
	Balance float64
}
