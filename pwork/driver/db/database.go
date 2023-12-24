package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func SqlConnect() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DSN"), //"host=pg_db user=root password=root dbname=location port=5432 sslmode=disable"
		// DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
}

func InitDB() {
	result := DB.Create(Location{})
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	result = DB.Create(Location{})
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
}
