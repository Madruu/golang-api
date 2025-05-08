package models

type Client struct {
	//gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	Name      string
	Email     *string
	Password  string
	BankRefer int  `json:"bank_id"`
	Bank      Bank `gorm:"foreignKey:BankRefer;references:ID"` //referencing Bank model
	//Bank     Bank `gorm:"embedded"`
	//Bank string
	Age        uint8
	GamesRefer int  `json:"game_id"`
	Games      Game `gorm:"foreignKey:GamesRefer;references:ID"` //Referencing Game model
	//Games Game `gorm:"embedded"`
	//Games string
}
