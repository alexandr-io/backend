package database

import "go.mongodb.org/mongo-driver/mongo"

// CollectionUser is the name of the user collection in mongodb
const CollectionUser = "user"

// CollectionFeedback is the name of the feedback collection in mongodb
const CollectionFeedback = "feedback"

// UserCollection collection for user
var UserCollection *mongo.Collection

// FeedbackCollection collection for feedback
var FeedbackCollection *mongo.Collection
