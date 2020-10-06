package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InstanceData contains the Mongo client and database objects.
type InstanceData struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// Instance is the instanceMongoData connected to the database.
// It should be used to interact with the database.
var Instance InstanceData

// Database settings
const dbName = "user"

var mongoURI = fmt.Sprintf(
	"mongodb+srv://%s:%s@%s/%s?authSource=admin&readPreference=primary&appname=UserService&ssl=false",
	os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
	os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	os.Getenv("MONGO_URL"),
	dbName)

// ConnectToMongo is connecting the service to mongodb using mongoURI.
// After success, instanceMongo is filled with the db client and db handler.
func ConnectToMongo() {
	if _, ok := os.LookupEnv("DEV"); ok {
		mongoURI = fmt.Sprintf(
			"mongodb://%s:%s@%s:27017/%s?authSource=admin&readPreference=primary&appname=UserService&ssl=false",
			os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
			os.Getenv("MONGO_URL"),
			dbName)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)

	Instance = InstanceData{
		Client: client,
		Db:     db,
	}
}

// InitCollections call the functions that init the collections.
func InitCollections() {
	createUserUniqueIndexes()
}
