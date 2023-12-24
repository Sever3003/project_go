package db

type Location struct {
	Id         uint `gorm:"primaryKey"`
	Locations  string
	DriverUuid string
}
