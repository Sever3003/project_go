package db

import (
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	Id             bson.ObjectId `bson:"id"`
	Clientname     string        `bson:"clientname"`
	Location_Start string        `bson:"location_start"`
	Location_End   string        `bson:"location_end"`
	Status         int           `bson:"status"`
}

type DriverInfo struct {
	Id         bson.ObjectId `bson:"id"`
	DriverName string        `bson:"drivername"`
	Uuid       uuid.UUID     `bson:"uuid"`
}
