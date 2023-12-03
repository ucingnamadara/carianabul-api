package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type PostgresImpl struct {
	DB *gorm.DB
}

func NewPostrgresClient() *PostgresImpl {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Can't connect db. Err: %s", err)
	}

	return &PostgresImpl{DB: db}

}
