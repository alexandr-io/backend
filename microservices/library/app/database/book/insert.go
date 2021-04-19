package book

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert on the database a new book in a library.
func Insert(bookData data.Book) (*data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Instance.Db.Collection(database.CollectionBook)

	result, err := collection.InsertOne(ctx, bookData)
	if err != nil {
		return nil, err
	}

	bookData.ID = result.InsertedID.(primitive.ObjectID)
	return &bookData, nil
}
