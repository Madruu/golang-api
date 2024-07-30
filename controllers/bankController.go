package controllers

import (
	"github.com/Madruu/golangDatabase/models"
	"github.com/gin-gonic/gin"
)

// Bank serializer
type Bank struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Number  string  `json:"number"`
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

func CreateResponseBank(bankModel models.Bank) Bank {
	return Bank{ID: bankModel.ID, Name: bankModel.Name,
		Number: bankModel.Number, UserID: bankModel.UserID, Balance: bankModel.Balance}
}

func CreateBank(c *gin.Context) {

}
