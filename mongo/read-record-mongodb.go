package mongo

import (
	"context"
	"fmt"
)

func ReadRecord() {
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}

	err = session.StartTransaction()
	if err != nil {
		panic(err)
	}

	database := client.Database("test")
	collection := database.Collection("data")

	var filter interface{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		row, err := cursor.DecodeBytes()
		if err != nil {
			panic(err)
		}
		fmt.Println(row.String())
	}

}
