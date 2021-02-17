package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func initInvitationCollection() {
	invitationCollection := Instance.Db.Collection(CollectionInvitation)
	_, err := invitationCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"token", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

// InitCollections call the functions that init the collections.
func InitCollections() {
	initInvitationCollection()
}
