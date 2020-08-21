package main

import (
	"github.com/alexandr-io/backend_errors"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

type user struct {
	ID       string `json:"-" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"-" bson:"password,omitempty"`
}

// getOneUserByID
func getOneUserByID(ctx *fiber.Ctx, id interface{}) (*user, bool) {
	collection := instanceMongo.Db.Collection(collectionUser)
	filter := bson.D{{Key: "_id", Value: id}}
	object := &user{}

	filteredSingleResult := collection.FindOne(ctx.Fasthttp, filter)
	err := filteredSingleResult.Decode(object)
	if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return nil, false
	}
	return object, true
}
