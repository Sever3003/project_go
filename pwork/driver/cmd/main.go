package main

import (
	"log"

	"github.com/Gwoop/Driver/api"
	"github.com/Gwoop/Driver/db"
	"github.com/joho/godotenv"
)

func main() {
	Init()
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла!" + err.Error())
	}
	db.SqlConnect()
	db.Migrate()
	api.API()
}
