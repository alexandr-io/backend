package feedback

import (
	"context"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert inserts a new feedback in the database
func Insert(feedback data.Feedback) (*data.Feedback, error) {
	insertedResult, err := database.FeedbackCollection.InsertOne(context.Background(), feedback)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	feedback.ID = insertedResult.InsertedID.(primitive.ObjectID)

	return &feedback, nil
}
