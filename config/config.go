package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DbInstance()

func DbInstance() *mongo.Client {

	uri := "mongodb://localhost:27017/caloriesdb"
	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	log.Fatal("Set your 'Mongo_URI' environment variables." + "See: " +
	// 		"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	// }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongodb")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("entry").Collection(collectionName)
	return collection
}
