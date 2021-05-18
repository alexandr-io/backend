package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthInfo struct {
	ID       primitive.ObjectID
	Username string
	Email    string
}
