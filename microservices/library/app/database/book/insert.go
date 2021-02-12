package book

import (
	"context"
	"github.com/alexandr-io/backend/library/database"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert on the database a new book in a library.
func Insert(DBBook data.BookData) (*data.BookData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(database.CollectionBook)

	result, err := collection.InsertOne(ctx, DBBook)
	if err != nil {
		return nil, err
	}
	DBBook.ID = result.InsertedID.(primitive.ObjectID)
	return &DBBook, nil
}
