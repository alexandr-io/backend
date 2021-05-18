package database

import "go.mongodb.org/mongo-driver/mongo"

// CollectionBook is the name of the book collection in mongodb
const CollectionBook = "book"

// BookCollection collection for book
var BookCollection *mongo.Collection
