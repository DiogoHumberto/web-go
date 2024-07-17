package db

import (
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"study.com/golang-web/models"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {
	// Connect to your postgres DB.
	stringDeConexao := "host=localhost port=5432 user=docker password=docker dbname=api-produto sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Product{}, &models.User{})
}
