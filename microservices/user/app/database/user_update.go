package database

import (
	"context"
	"log"

	"github.com/alexandr-io/backend/user/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateUser take a data.User and an ID to update a user in database
func UpdateUser(id string, user data.User) (data.User, error) {
	// Transform ID to bson id object
	userObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return data.User{}, err
	}

	// Update data
	userCollection := Instance.Db.Collection(CollectionUser)
	if _, err = userCollection.UpdateOne(
		context.Background(),
		bson.D{
			{"_id", userObjectID},
		},
		bson.D{
			{"$set", user},
		},
	); err != nil {
		log.Println(err)
		return data.User{}, err
	}

	// Retrieve updated data
	updatedUser, err := GetUserByID(userObjectID)
	if err != nil {
		return data.User{}, err
	}
	return *updatedUser, nil
}
