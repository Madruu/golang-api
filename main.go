package main

import (
	"github.com/Madruu/golangDatabase/controllers"
	"github.com/Madruu/golangDatabase/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	//Game Routes
	r.POST("/game", controllers.CreateGame)
	r.GET("/games", controllers.GetGames)
	r.GET("/games/:id", controllers.GetGameById)
	r.PUT("/games/:id", controllers.UpdateGame)
	r.DELETE("/games/:id", controllers.DeleteGame)

	//Bank routes
	r.POST("/bank", controllers.CreateBank)
	r.GET("/banks", controllers.GetBanks)
	r.GET("/banks/:id", controllers.GetBankById)

	//Client routes
	r.POST("/client", controllers.RegisterClient)

	r.Run() //Simple example on how to run a server using gin
}
