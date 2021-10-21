package server

import (
	"context"
	"go-mongo/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	dataBase   string
	collection string
}

var ctx = context.Background()

func NewCollection(db string, col string) *Collection {
	return &Collection{
		dataBase:   db,
		collection: col,
	}
}

func (c *Collection) getCollection() *mongo.Collection {
	return database.Db.Database(c.dataBase).Collection(c.collection)
}

func (c *Collection) Insert(data map[string]string) {
	col := c.getCollection()
	col.InsertOne(ctx, data)
}
