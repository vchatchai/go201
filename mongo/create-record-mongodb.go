package mongo

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo/readpref"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func init() {
	var err error
	client, err = mongo.NewClient(MONGO_DB_URL)
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
}

func OneRecord() {
	database := client.Database("test")
	collection := database.Collection("data")

	res, err := collection.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		panic(err)
	}

	id := res.InsertedID
	fmt.Println(id)
}
