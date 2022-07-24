package routes

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBInstance()

func DBInstance() *mongo.Client {
	mongoURL := "mongodb://localhost:49153"

	// setup options for MongoDB connection
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "docker",
		Password: "mongopw",
	})

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to MongoDB")
	return conn
}

func OpenCollection(client *mongo.Client, name string) *mongo.Collection {
	return client.Database("caloriesdb").Collection(name)
}
