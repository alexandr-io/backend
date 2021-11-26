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
const dbName = "library"

// ConnectToMongo is connecting the service to mongodb using mongoURI.
// After success, instanceMongo is filled with the db client and db handler.
func ConnectToMongo() {
	mongoURI := fmt.Sprintf(
		"mongodb://%s:%s@%s:27017/%s?authSource=admin&readPreference=primary&appname=LibraryService&ssl=false",
		os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		os.Getenv("MONGO_URL"),
		dbName)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println(mongoURI)
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)

	Instance = InstanceData{
		Client: client,
		Db:     db,
	}

	BookDB = NewBookCollection(db)
	UserLibraryDB = NewUserLibraryCollection(db)
	LibraryDB = NewLibraryCollection(db)
	BookProgressDB = NewBookProgressCollection(db)
	GroupDB = NewGroupCollection(db)
	UserDataDB = NewUserDataCollection(db)
	ProgressSpeedDB = NewProgressSpeedCollection(db)
}
