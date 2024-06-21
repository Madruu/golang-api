package controllers

import (
	"net/http"

	"github.com/Madruu/golangDatabase/initializers"
	"github.com/Madruu/golangDatabase/models"
	"github.com/gin-gonic/gin"
)

func CreateGame(c *gin.Context) {

	//Get data off req.body
	var body struct {
		Name        string
		Price       float64
		Platform    string
		ReleaseDate string
	}

	//Adding to model
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return //Exit
	}

	//Creating game
	game := models.Game{
		Name:        body.Name,
		Price:       body.Price,
		Platform:    body.Platform,
		ReleaseDate: body.ReleaseDate,
	}

	result := initializers.DB.Create(&game)

	//Result can return an error (according to GORM)
	if result.Error != nil {
		c.Status(400)
		return //Exit
	}

	c.JSON(200, gin.H{
		"game": game,
	})

}

func GetGame(c *gin.Context) {
	var games []models.Game //Creating slice of games

	result := initializers.DB.Find(&games)

	if result.Error != nil {
		c.Status(400)
		return //Exit
	}

	c.JSON(200, gin.H{
		"game": games,
	})

}

func DeleteGame(c *gin.Context) {
	//Getting id from URL
	id := c.Param("id")

	//Find game by id
	initializers.DB.Delete(&models.Game{}, id) //What to delete, using id as parameter

	c.JSON(200, gin.H{
		"message": "Game deleted successfuly",
	})
}

func GetGameById(c *gin.Context) {
	//Getting if from URL
	id := c.Param("id")

	//Storing game
	var game models.Game

	//Getting the game
	initializers.DB.First(&game, id)

	c.JSON(200, gin.H{
		"game": game,
	})
}
