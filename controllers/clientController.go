package controllers

import (
	"log"
	"net/http"

	"github.com/Madruu/golangDatabase/initializers"
	"github.com/Madruu/golangDatabase/models"
	"github.com/Madruu/golangDatabase/services"
	"github.com/gin-gonic/gin"
)

// Struct serializer
type Client struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
	Bank     Bank    `json:"bank"`
	Age      uint8   `json:"age"`
	Games    Game    `json:"games"`
}

func CreateResponseClient(client models.Client, bank Bank, game Game) Client {
	return Client{ID: client.ID, Name: client.Name, Email: client.Email,
		Password: client.Password, Bank: bank, Age: client.Age, Games: game}
}

func RegisterClient(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBind(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Check if the associated bank  exists
	var bank models.Bank
	if err := findBank(client.BankRefer, &bank); err != nil {
		c.Status(400)
		return
	}

	//Check if the associated game  exists
	var game models.Game
	if err := findGame(client.GamesRefer, &game); err != nil {
		c.Status(400)
		return
	}

	hashedPassword, err := services.HashPasswordFromClient(client.Password)

	client.Password = hashedPassword

	initializers.DB.Create(&client)

	if err != nil {
		log.Fatal("Couldn't hash password")
	}

	responseBank := CreateResponseBank(bank)
	responseGame := CreateResponseGame(game)
	responseClient := CreateResponseClient(client, responseBank, responseGame)

	c.JSON(200, gin.H{
		"client": responseClient,
	})
}
