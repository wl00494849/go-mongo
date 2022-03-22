package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client

func init() {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	Db = client
}
