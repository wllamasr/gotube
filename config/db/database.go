package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wllamasr/golangtube/models"
	"os"
)

var (
	Client *gorm.DB
)

/*
Create the SQL Client.
 */
func LoadClient() {
	var (
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		schema   = os.Getenv("DB_DATABASE")
		port     = os.Getenv("DB_PORT")
	)
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, schema)

	var err error

	Client, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	Client.LogMode(true)
	migrateDatabase()
}

func migrateDatabase() {
	fmt.Println("Migrating")
	Client.AutoMigrate(&models.User{}, &models.Upload{}, &models.Comment{})
}
