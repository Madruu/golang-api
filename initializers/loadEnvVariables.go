package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao rodar arquivo .env")
	}
}
