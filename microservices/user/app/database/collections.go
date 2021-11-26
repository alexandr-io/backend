package database

import "go.mongodb.org/mongo-driver/mongo"

// CollectionUser is the name of the user collection in mongodb
const CollectionUser = "user"

// UserCollection collection for user
var UserCollection *mongo.Collection
