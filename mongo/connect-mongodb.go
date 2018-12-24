package mongo

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	MONGO_DB_URL = "mongodb://root:manager1@localhost:27017" // "mongodb://root:manager1@127.0.0.1:27017"
)

var client *mongo.Client

func init() {

}

func ConnectMongoDB() {
	fmt.Println("Init")
	client, err := mongo.NewClient(MONGO_DB_URL)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	var filter interface{}
	result, err := client.ListDatabases(ctx, filter, &options.ListDatabasesOptions{})

	if err != nil {
		panic(err)
	}

	for _, database := range result.Databases {
		fmt.Printf(database.Name)

		collections, err := client.Database(database.Name).ListCollections(ctx, filter)

		if err != nil {
			panic(err)
		}

		for collections.Next(ctx) {
			row, err := collections.DecodeBytes()
			if err != nil {
				panic(err)
			}
			fmt.Println(row.String())
		}

		collections.Close(ctx)

	}
	fmt.Println("Inited")

	client.Disconnect(ctx)
}
