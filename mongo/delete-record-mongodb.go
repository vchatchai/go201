package mongo

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
)

func Delete() {
	database := client.Database("test")
	collection := database.Collection("data")

	result, err := collection.DeleteMany(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Deleted %v", result.DeletedCount)
}
