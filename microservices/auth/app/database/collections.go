package database

import "go.mongodb.org/mongo-driver/mongo"

// CollectionInvitation is the name of the invitation collection in the database
const CollectionInvitation = "invitation"

// InvitationCollection collection for invitation
var InvitationCollection *mongo.Collection
