package models

type Game struct {
	//gorm.Model
	ID          uint `json:"id" gorm:"primaryKey"`
	Name        string
	Price       float64
	Platform    string
	ReleaseDate string
}
