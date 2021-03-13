package user

import (
	"context"
	"errors"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

// Insert insert a new user into the database.
func Insert(user data.User) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := database.Instance.Db.Collection(database.CollectionUser)

	insertedResult, err := userCollection.InsertOne(ctx, user)
	if database.IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		err := checkRegisterFieldDuplication(user)
		var badInput *data.BadInputError
		if errors.As(err, &badInput) {
			return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, badInput.Error())
		}
	} else if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	user.ID = insertedResult.InsertedID.(primitive.ObjectID).Hex()
	return &user, nil
}
