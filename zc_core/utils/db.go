package utils

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


//cluster_url := "mongodb://zuri:<password>@127.0.0.1:27017/myFirstDatabase?authSource=admin"


// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
//getMongoDbConnection get connection of mongodb
func getMongoDbConnection() (*mongo.Client, error) {
	//"mongodb+srv://zuri:<password>@cluster0.hepte.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	cluster_url := "mongodb+srv://zuri:<password>@cluster0.hepte.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(cluster_url)

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
func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := getMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
