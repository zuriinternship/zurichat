package utils

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
//getMongoDbConnection get connection of mongodb
func getMongoDbConnection() (*mongo.Client, error) {
	config, err := LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}

	clientOptions := options.Client().ApplyURI(config.DBHost)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

//GetMongoDbCollection get collection inside your db, this function can be exorted
func GetMongoDbCollection(CollectionName string) (*mongo.Collection, error) {
	config, err := LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}

	client, err := getMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DBDatabase).Collection(CollectionName)

	return collection, nil
}
