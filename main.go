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

	r.POST("/clients", controllers.CreateClient)
	r.POST("/bank", controllers.CreateBank)
	r.DELETE("/clients/:id", controllers.DeleteClient)
	r.GET("/clients", controllers.GetClients)
	r.GET("/clients/:id", controllers.GetClientById)

	r.Run() //Simple example on how to run a server using gin
}
