package main

import (
	"github.com/Madruu/golangDatabase/initializers"
	"github.com/Madruu/golangDatabase/models" // Add this import statement
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Client{})
	initializers.DB.AutoMigrate(&models.Bank{})
}
