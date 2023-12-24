package main

import (
	"log"

	"driverserv/api"
	"driverserv/db"

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
	api.API()
}
