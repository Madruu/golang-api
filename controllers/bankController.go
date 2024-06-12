package controllers

import (
	"github.com/Madruu/golangDatabase/initializers"
	"github.com/Madruu/golangDatabase/models"
	"github.com/gin-gonic/gin"
)

func CreateBank(c *gin.Context) {
	var body struct {
		Name    string
		Number  string
		UserID  uint
		Balance float64
	}

	c.Bind(&body)

	bank := models.Bank{Name: body.Name, Number: body.Number, UserID: body.UserID, Balance: body.Balance}

	result := initializers.DB.Create(&bank)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"bank": bank,
	})
}
