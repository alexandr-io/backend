package data

import (
	"github.com/Alexandr-io/Backend/User/database"
	"github.com/alexandr-io/backend_errors"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// User defines the structure for an API user
// swagger:model
type User struct {
	ID string `json:"-" bson:"_id,omitempty"`
	// The username of this user
	// required: true
	// example: john
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	// The email address of this user
	// required: true
	// example: john@provider.net
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"-" bson:"password,omitempty"`
	// The authentication token of this user
	// example: eyJhb[...]FYqf4
	JWT string `json:"jwt" bson:"-"`
}

// GetUserByID get an user by it's given id.
// In case of error, the proper error is set to the context and false is returned.
func GetUserByID(ctx *fiber.Ctx, id interface{}) (*User, bool) {
	collection := database.Instance.Db.Collection(database.CollectionUser)
	filter := bson.D{{Key: "_id", Value: id}}
	object := &User{}

	filteredSingleResult := collection.FindOne(ctx.Fasthttp, filter)
	err := filteredSingleResult.Decode(object)
	if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return nil, false
	}
	return object, true
}
