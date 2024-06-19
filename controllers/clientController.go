package controllers

import (
	"log"
	"net/http"

	"github.com/Madruu/golangDatabase/initializers"
	"github.com/Madruu/golangDatabase/models"
	"github.com/Madruu/golangDatabase/services"
	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	//Get data off req.body
	var body struct {
		Name     string
		Email    *string
		Password string
		Bank     models.Bank `json:"bank"`
		//Bank  string
		Age uint8
		//Games string
		Games models.Game
	}

	//c.Bind(&body)

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := models.Client{Name: body.Name, Email: body.Email,
		Password: body.Password, Bank: body.Bank, Age: body.Age, Games: body.Games}

	//Hashing password
	hashedPassword, err := services.HashPasswordFromClient(client.Password)

	client.Password = hashedPassword

	if err != nil {
		log.Fatal("Couldnt hash password")
	}

	result := initializers.DB.Create(&client)

	//Result can return an error (gorm)
	if result.Error != nil {
		c.Status(400)
		return
	}

	initializers.DB.Preload("Bank").Preload("Games").Find(&client) //Preloading foreign keys

	c.JSON(200, gin.H{
		"client": client,
	})
}

func GetClients(c *gin.Context) {
	var clients []models.Client

	result := initializers.DB.Find(&clients)

	if result.Error != nil {
		c.Status(400)
		return
	}

	initializers.DB.Preload("Bank").Preload("Games").Find(&clients)

	c.JSON(200, gin.H{
		"client": clients,
	})
}

func DeleteClient(c *gin.Context) {

	//Getting id
	id := c.Param("id")

	//Delete the post from the id
	initializers.DB.Delete(&models.Client{}, id) //delete from the id

	//Respond
	c.JSON(200, gin.H{
		"message": "Client deleted",
	})

}
