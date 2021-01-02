package database

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CollectionInvitation is the invitation collection
const CollectionInvitation = "invitation"

//
// Getters
//

// GetInvitationByToken get an invitation by it's given token.
func GetInvitationByToken(token string) (*data.Invitation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionInvitation)
	filter := bson.D{{Key: "token", Value: token}}
	object := &data.Invitation{}

	filteredSingleResult := collection.FindOne(ctx, filter)
	err := filteredSingleResult.Decode(object)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return object, nil
}

//
// Setters
//

// InsertInvitation insert a new invitation into the database.
func InsertInvitation(invitation data.Invitation) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	invitationCollection := Instance.Db.Collection(CollectionInvitation)

	insertedResult, err := invitationCollection.InsertOne(ctx, data.Invitation{
		Token:  invitation.Token,
		Used:   nil,
		UserID: nil,
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return insertedResult, nil
}

// UpdateInvitation take a data.Invitation to update an invitation in database
func UpdateInvitation(invitation data.Invitation) (data.Invitation, error) {
	// Update data
	invitationCollection := Instance.Db.Collection(CollectionInvitation)
	if _, err := invitationCollection.UpdateOne(
		context.Background(),
		bson.D{
			{"token", invitation.Token},
		},
		bson.D{
			{"$set", invitation},
		},
	); err != nil {
		return data.Invitation{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	// Retrieve updated data
	updatedInvitation, err := GetInvitationByToken(invitation.Token)
	if err != nil {
		return data.Invitation{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return *updatedInvitation, nil
}

// DeleteInvitation delete an invitation corresponding to the given invitation token
func DeleteInvitation(token string) error {
	invitationCollection := Instance.Db.Collection(CollectionInvitation)

	result, err := invitationCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "token", Value: token},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "can't find invitation to delete")
	}
	return nil
}
