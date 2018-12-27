package mongo

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
)

func Update() {
	database := client.Database("test")
	collection := database.Collection("data")

	filter := bson.M{"name": "pi"}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		raw := bson.M{}
		// raw, err := cursor.DecodeBytes()
		cursor.Decode(&raw)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v", raw)
		collection.UpdateOne(context.TODO(), filter, raw)
	}

}
