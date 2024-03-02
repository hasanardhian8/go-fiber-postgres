package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hasanardhian8/go-fiber-postgres/models"
	"github.com/joho/godotenv"
)

var DBConn *gorm.DB

func DatabaseConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	DBConn, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println("disini error")
		panic(err.Error())
	}
	fmt.Println("connecting to database.....")

	DBConn.AutoMigrate(
		&models.Penduduks{},
		&models.Bayars{},
		&models.Petugases{},
	)
}
