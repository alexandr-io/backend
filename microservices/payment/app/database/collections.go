package database

import "go.mongodb.org/mongo-driver/mongo"

// CollectionSubscriptions is the name of the subscriptions collection in the database
const CollectionSubscriptions = "subscriptions"

// CollectionCustomer is the name of the user's subscription collection in the database
const CollectionCustomer = "customer"

// SubscriptionsCollection collection for subscriptions
var SubscriptionsCollection *mongo.Collection

// CustomerCollection collection for user's subscription
var CustomerCollection *mongo.Collection