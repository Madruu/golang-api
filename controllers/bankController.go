package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Madruu/golangDatabase/initializers"
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
	var bank models.Bank

	if err := c.ShouldBind(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	initializers.DB.Create(&bank)
	responseBank := CreateResponseBank(bank)

	c.JSON(200, gin.H{
		"bank": responseBank,
	})
}

func GetBanks(c *gin.Context) {
	var banks []models.Bank

	result := initializers.DB.Find(&banks)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"banks": banks,
	})
}

func findBank(id int, bank *models.Bank) error {
	initializers.DB.First(bank, "id = ?", id)

	//Check if bank exists
	if bank.ID == 0 {
		return errors.New("Bank not found")
	}

	return nil
}

func GetBankById(c *gin.Context) {
	var bank models.Bank

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := findBank(id, &bank); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	responseBank := CreateResponseBank(bank)

	c.JSON(200, gin.H{
		"bank": responseBank,
	})

}
