package utils

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
//getMongoDbConnection get connection of mongodb
func getMongoDbConnection() (*mongo.Client, context.Context, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientOptions := options.Client().ApplyURI(config.DBHost)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, nil
}

//GetMongoDbCollection get collection inside your db, this function can be exorted
func GetMongoDbCollection(CollectionName string) (*mongo.Collection, context.Context, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, ctx, err := getMongoDbConnection()

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(config.DBDatabase).Collection(CollectionName)

	return collection, ctx, nil
}

// get MongoDb documents for a collection
func GetMongoDbDocs(CollectionName string, filter map[string]interface{}) ([]bson.M, error) {
	collection, ctx, err := GetMongoDbCollection(CollectionName)
	if err != nil {
		return nil, err
	}

	var data []bson.M
	filterCursor, err := collection.Find(ctx, MapToBson(filter))
	if err != nil {
		return nil, err
	}
	if err = filterCursor.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// get single MongoDb document for a collection
func GetMongoDbDoc(CollectionName string, filter map[string]interface{}) (bson.M, error) {
	collection, ctx, err := GetMongoDbCollection(CollectionName)
	if err != nil {
		return nil, err
	}

	var data bson.M
	if err = collection.FindOne(ctx, MapToBson(filter)).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func CreateMongoDbDoc(CollectionName string, data map[string]interface{}) (*mongo.InsertOneResult, error) {
	collection, ctx, err := GetMongoDbCollection(CollectionName)
	if err != nil {
		return nil, err
	}
	res, err := collection.InsertOne(ctx, MapToBson(data))

	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}
