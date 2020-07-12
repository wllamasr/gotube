package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wllamasr/golangtube/models"
	"os"
)

var (
	Client *gorm.DB
)

func LoadClient() {
	var (
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		schema   = os.Getenv("DB_DATABASE")
		port     = os.Getenv("DB_PORT")
	)

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, schema, password)

	var error error

	Client, error = gorm.Open("postgres", dataSourceName)

	if error != nil {
		fmt.Println(error.Error())
		panic(error.Error())
	}

	migrateDatabase()
}

func migrateDatabase() {
	Client.AutoMigrate(&models.User{})
}
