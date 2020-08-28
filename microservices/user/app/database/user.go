package database

import (
	"log"
	"net/http"

	"github.com/Alexandr-io/Backend/User/data"
	"github.com/alexandr-io/backend_errors"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CollectionUser is the name of the user collection in mongodb
const CollectionUser = "user"

//
// Getters
//

// FindOneWithFilter fill the given object with a mongodb single result filtered by the given filters.
func FindOneWithFilter(ctx *fiber.Ctx, object interface{}, filters interface{}) error {
	collection := Instance.Db.Collection(CollectionUser)
	filteredSingleResult := collection.FindOne(ctx.Fasthttp, filters)
	return filteredSingleResult.Decode(object)
}

// GetUserByID get an user by it's given id.
// In case of error, the proper error is set to the context and false is returned.
func GetUserByID(ctx *fiber.Ctx, id interface{}) (*data.User, bool) {
	collection := Instance.Db.Collection(CollectionUser)
	filter := bson.D{{Key: "_id", Value: id}}
	object := &data.User{}

	filteredSingleResult := collection.FindOne(ctx.Fasthttp, filter)
	err := filteredSingleResult.Decode(object)
	if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return nil, false
	}
	return object, true
}

// GetUserByLogin get an user by it's given login (username or email).
// In case of error, the proper error is set to the context and false is returned.
func GetUserByLogin(ctx *fiber.Ctx, login string) (*data.User, bool) {
	collection := Instance.Db.Collection(CollectionUser)
	object := &data.User{}

	// Get user by username
	usernameFilter := bson.D{{Key: "username", Value: login}}
	filteredByUsernameSingleResult := collection.FindOne(ctx.Fasthttp, usernameFilter)
	// Return the user object if user is found
	if err := filteredByUsernameSingleResult.Decode(object); err == nil {
		return object, true
	}

	// Get user by email
	emailFilter := bson.D{{Key: "email", Value: login}}
	filteredByEmailSingleResult := collection.FindOne(ctx.Fasthttp, emailFilter)
	// Return a login error if the user is not found
	if err := filteredByEmailSingleResult.Decode(object); err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest).SendBytes(
			backend_errors.BadInputJSONFromType("login", string(backend_errors.Login)))
		return nil, false
	}

	// Return the email user object
	return object, true
}

//
// Setters
//

// InsertUserRegister insert a new user into the database.
// The proper error is set to the context in case of error or duplication.
func InsertUserRegister(ctx *fiber.Ctx, user data.User) *mongo.InsertOneResult {
	userCollection := Instance.Db.Collection(CollectionUser)

	insertedResult, err := userCollection.InsertOne(ctx.Fasthttp, data.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		checkRegisterFieldDuplication(ctx, user)
		return nil
	} else if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return nil
	}
	return insertedResult
}

// checkRegisterFieldDuplication check which field is a duplication on a register call.
// The correct http error and content is handled and returned.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
func checkRegisterFieldDuplication(ctx *fiber.Ctx, user data.User) {
	errorsFields := make(map[string]string)

	// Check if the duplication is for the email field
	filter := bson.D{{Key: "email", Value: user.Email}}
	filteredByEmailUser := &data.User{}
	err := FindOneWithFilter(ctx, filteredByEmailUser, filter)
	if err == nil && filteredByEmailUser.Email == user.Email {
		errorsFields["email"] = "Email has already been taken."
	} else if err != nil {
		log.Println(err)
	}

	// Check if the duplication is for the username field
	filter = bson.D{{Key: "username", Value: user.Username}}
	filteredByUsernameUser := &data.User{}
	err = FindOneWithFilter(ctx, filteredByUsernameUser, filter)
	if err == nil && filteredByUsernameUser.Username == user.Username {
		errorsFields["username"] = "Username has already been taken."
	} else if err != nil {
		log.Println(err)
	}

	ctx.Status(http.StatusBadRequest).SendBytes(backend_errors.BadInputsJSON(errorsFields))
}
