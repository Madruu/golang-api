package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordFromClient(client string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(client), 14)

	if err != nil {
		log.Fatal("Erro ao gerar hash da senha")
	}

	return string(bytes), nil
}
