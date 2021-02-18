package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mrogerius/educa/domain"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro ao carregar")
	}
	//pacote do GO que le variavel de ambiente. Esta no arquivo .ENV

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting DBase Fail: %v", err)
		panic(err)
	}
	//Fecha a conexao
	//defer db.Close()
	//cria a tabela conforme o definido em User
	db.AutoMigrate(&domain.User{})

	return db
}
