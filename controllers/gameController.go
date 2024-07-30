package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Madruu/golangDatabase/initializers"
	"github.com/Madruu/golangDatabase/models"
	"github.com/gin-gonic/gin"
)

// Not the model. This is the serializer
type Game struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Platform    string  `json:"platform"`
	ReleaseDate string  `json:"release_date"`
}

func CreateResponseGame(gameModel models.Game) Game {
	return Game{ID: gameModel.ID, Name: gameModel.Name,
		Price: gameModel.Price, Platform: gameModel.Platform, ReleaseDate: gameModel.ReleaseDate}
}

func CreateGame(c *gin.Context) {
	var game models.Game

	if err := c.ShouldBind(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	initializers.DB.Create(&game)
	responseGame := CreateResponseGame(game)

	c.JSON(200, gin.H{
		"game": responseGame,
	})
}

func GetGames(c *gin.Context) {
	var games []models.Game

	result := initializers.DB.Find(&games)

	if result.Error != nil {
		c.Status(400)
		return //Exit
	}

	c.JSON(200, gin.H{
		"games": games,
	})
}

func findGame(id int, game *models.Game) error {
	initializers.DB.First(&game, "id = ?", id)

	//Check if Game exists (if id equals 0, doesn't exist)
	if game.ID == 0 {
		return errors.New("game not found")
	}
	return nil
}

func GetGameById(c *gin.Context) {
	var game models.Game

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := findGame(id, &game); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	responseGame := CreateResponseGame(game)

	c.JSON(200, gin.H{
		"game": responseGame,
	})
}

func UpdateGame(c *gin.Context) {
	var game models.Game

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := findGame(id, &game); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var updateGame struct {
		Name        string
		Price       float64
		Platform    string
		ReleaseDate string
	}

	if err := c.ShouldBind(&updateGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.First(&game, id)

	initializers.DB.Model(&game).Updates(models.Game{
		Name:        updateGame.Name,
		Price:       updateGame.Price,
		Platform:    updateGame.Platform,
		ReleaseDate: updateGame.ReleaseDate,
	})

	c.JSON(200, gin.H{
		"game": game,
	})
}

func DeleteGame(c *gin.Context) {
	var game models.Game

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := findGame(id, &game); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Delete(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Game deleted",
	})
}
