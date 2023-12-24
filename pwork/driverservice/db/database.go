package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func SqlConnect() {
	// Создание клиета для подключения

	var dbUser string = "root"
	var dbPassword string = "local"
	dbURL := os.Getenv("MONGO_URL")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPassword + "@" + dbURL + ":27017/?connect=direct")) //"mongodb://localhost:27017"
	if err != nil {
		log.Fatal(err)
	}

	// создание соединения
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// проверка соединения
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Упешное соединение с MongoDB!")

	client.Database("drivers").Collection("order")
	client.Database("drivers").Collection("driverinfo")

	InitDataOrder()
	InitDataDriver()
}

func InitDataOrder() {
	collection := client.Database("drivers").Collection("order")

	//тестовые данные
	a := Order{Id: bson.NewObjectId(), Clientname: "Сафронов Вадим", Location_Start: "55.408442, 37.563368", Location_End: "55.431094, 37.566629", Status: 0}
	b := Order{Id: bson.NewObjectId(), Clientname: "Тестеров Вадим", Location_Start: "55.408442, 37.563368", Location_End: "55.431094, 37.566629", Status: 0}
	i := []interface{}{a, b}
	insertManyResult, err := collection.InsertMany(context.TODO(), i)
	if err != nil {
		fmt.Print("dfgdfgdfg")
		log.Fatal(err)

	}
	fmt.Println("Добавление тестовых документов: ", insertManyResult.InsertedIDs)
}

func InitDataDriver() {

	d := DriverInfo{DriverName: "Рамзан кадыров"}
	err := AddDriver(d)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//тестовые данные

	fmt.Println("Добавление тестовых водителей")
}

type Location struct {
	Locations  string `json:"location"`
	DriverUuid string `json:"uuid"`
}

func AddDriver(d DriverInfo) error {
	collection := client.Database("drivers").Collection("driverinfo")
	d.Id = bson.NewObjectId()
	d.Uuid = uuid.New()
	_, err := collection.InsertOne(context.TODO(), d)
	if err != nil {
		log.Fatal(err)
		return err
	}
	l := &Location{Locations: "", DriverUuid: d.Uuid.String()}

	data, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("LOCATION"), bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func GetDriver() ([]DriverInfo, error) {
	collection := client.Database("drivers").Collection("driverinfo")
	var result []DriverInfo

	filter := bson.M{}

	cur, err := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {

		var elem DriverInfo
		err := cur.Decode(&elem)
		if err != nil {
			return []DriverInfo{}, err
		}
		result = append(result, elem)
	}

	if err != nil {
		log.Fatal(err)
	}

	return result, nil

}

func AddDocument(o Order) error {
	collection := client.Database("drivers").Collection("order")
	o.Id = bson.NewObjectId()
	_, err := collection.InsertOne(context.TODO(), o)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func DeleteDocument(o Order) error {
	collection := client.Database("drivers").Collection("order")
	filter := bson.M{"clientname": o.Clientname}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func UpdateDocument(o Order) error {
	collection := client.Database("drivers").Collection("order")

	filter := bson.M{"clientname": o.Clientname}
	// bson.M{"model": "iPhone 8"}, bson.M{"$set":bson.M{"price":45000}}
	update := bson.M{"$set": bson.M{"status": o.Status}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}

func GetDocument() ([]Order, error) {
	collection := client.Database("drivers").Collection("order")
	var result []Order

	filter := bson.M{}

	cur, err := collection.Find(context.TODO(), filter)

	for cur.Next(context.TODO()) {

		var elem Order
		err := cur.Decode(&elem)
		if err != nil {
			return []Order{}, err
		}
		result = append(result, elem)
	}

	if err != nil {
		log.Fatal(err)
	}

	return result, nil

}
