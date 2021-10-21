package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client

func init() {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	Db = client
}
