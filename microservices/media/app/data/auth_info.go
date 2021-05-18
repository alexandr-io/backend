package data

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthInfo is used to retrieve the logged user data set by the auth middleware
type AuthInfo struct {
	ID       primitive.ObjectID
	Username string
	Email    string
}
