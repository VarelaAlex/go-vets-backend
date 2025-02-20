package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Ctx context.Context = context.TODO()
var DB *mongo.Database
var clientOptions = options.Client().ApplyURI("mongodb+srv://uo271288:aeRsGpiwwMz7p6Yg@cluster0.7xsq9.mongodb.net/")

func ConnectDB() {
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		DB = client.Database("go-vets-backend")
		log.Println("Successful connection")
	}
}
