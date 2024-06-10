package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Não utilize db(variavel local). crie variável global(DB)
var DB *gorm.DB

// Crie função para conexão com banco de dados
func ConnectToDB() {
	var err error
	//host=localhost port=5432 dbname=GolangDB user=postgres password=xxxxxxx connect_timeout=10 sslmode=prefer
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//tratamento de erro
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados")
	}
}
