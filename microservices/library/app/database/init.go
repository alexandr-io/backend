package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
)

// createLibraryUniqueIndexes init the library collection. It add a unique index to the username field.
func createLibraryUniqueIndexes() {
	libraryCollection := Instance.Db.Collection(CollectionLibrary)

	_, err := libraryCollection.Indexes().CreateOne(
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
