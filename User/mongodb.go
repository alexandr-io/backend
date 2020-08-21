package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// InstanceMongo contains the Mongo client and database objects.
type InstanceMongo struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// instanceMongo is the InstanceMongo connected to the database.
// It should be used to interact with the database.
var instanceMongo InstanceMongo

// Database settings
const dbName = "user"
const collectionUser = "user"

var mongoURI = fmt.Sprintf(
	"mongodb://%s:%s@%s:27017/%s?authSource=admin&readPreference=primary&appname=UserService&ssl=false",
	os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
	os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	os.Getenv("MONGO_URL"),
	dbName)

// connectToMongo is connecting the service to mongodb using mongoURI.
// After success, instanceMongo is filled with the db client and db handler.
func connectToMongo() {
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

	instanceMongo = InstanceMongo{
		Client: client,
		Db:     db,
	}
}

// initCollections call the functions that init the collections.
func initCollections() {
	createUserUniqueIndexes()
}

// createUserUniqueIndexes init the user collection. It add a unique index to the email and username field.
func createUserUniqueIndexes() {
	userCollection := instanceMongo.Db.Collection(collectionUser)
	_, err := userCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = userCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"username", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

// isMongoDupKey is checking whether is err given by a mongodb insertion is a duplication error.
// Possible error codes could be added if needed: https://jira.mongodb.org/browse/GODRIVER-972
func isMongoDupKey(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

// findOneWithFilter fill the given object with a mongodb single result filtered by the given filters.
func findOneWithFilter(ctx *fiber.Ctx, object interface{}, filters interface{}) error {
	collection := instanceMongo.Db.Collection(collectionUser)
	filteredSingleResult := collection.FindOne(ctx.Fasthttp, filters)
	return filteredSingleResult.Decode(object)
}
