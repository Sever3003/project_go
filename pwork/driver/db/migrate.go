package db

import "log"

//функция для осуществления миграций в бд pgsql
func Migrate() {
	err := DB.AutoMigrate(Location{})
	if err != nil {
		log.Println(err)
		return
	}
}
