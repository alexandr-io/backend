package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/berrors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CollectionUser is the name of the user collection in mongodb
const CollectionUser = "user"

//
// Getters
//

// FindOneWithFilter fill the given object with a mongodb single result filtered by the given filters.
func FindOneWithFilter(object interface{}, filters interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionUser)
	filteredSingleResult := collection.FindOne(ctx, filters)
	return filteredSingleResult.Decode(object)
}

// GetUserByID get an user by it's given id.
func GetUserByID(id interface{}) (*data.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionUser)
	filter := bson.D{{Key: "_id", Value: id}}
	object := &data.User{}

	filteredSingleResult := collection.FindOne(ctx, filter)
	err := filteredSingleResult.Decode(object)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return object, true
}

// GetUserByLogin get an user by it's given login (username or email).
// In case of error, the proper error is set to the context and false is returned.
func GetUserByLogin(login string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionUser)
	object := &data.User{}

	// Get user by username
	usernameFilter := bson.D{{Key: "username", Value: login}}
	filteredByUsernameSingleResult := collection.FindOne(ctx, usernameFilter)
	// Return the user object if user is found
	if err := filteredByUsernameSingleResult.Decode(object); err == nil {
		return object, err
	}

	// Get user by email
	emailFilter := bson.D{{Key: "email", Value: login}}
	filteredByEmailSingleResult := collection.FindOne(ctx, emailFilter)
	// Return a login error if the user is not found
	if err := filteredByEmailSingleResult.Decode(object); err != nil {
		log.Println(err)
		return nil, &data.BadInput{
			JSONError: berrors.BadInputJSONFromType("login", string(berrors.Login)),
			Err:       errors.New("can't find user with login " + login),
		}
	}

	// Return the email user object
	return object, nil
}

//
// Setters
//

// InsertUserRegister insert a new user into the database.
func InsertUserRegister(user data.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := Instance.Db.Collection(CollectionUser)

	insertedResult, err := userCollection.InsertOne(ctx, data.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		err := checkRegisterFieldDuplication(user)
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return insertedResult, nil
}

// checkRegisterFieldDuplication check which field is a duplication on a register call.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
// The error returned is a formatted json of berrors.BadInput
func checkRegisterFieldDuplication(user data.User) error {
	errorsFields := make(map[string]string)

	// Check if the duplication is for the email field
	filter := bson.D{{Key: "email", Value: user.Email}}
	filteredByEmailUser := &data.User{}
	err := FindOneWithFilter(filteredByEmailUser, filter)
	if err == nil && filteredByEmailUser.Email == user.Email {
		errorsFields["email"] = "Email has already been taken."
	} else if err != nil {
		log.Println(err)
		return err
	}

	// Check if the duplication is for the username field
	filter = bson.D{{Key: "username", Value: user.Username}}
	filteredByUsernameUser := &data.User{}
	err = FindOneWithFilter(filteredByUsernameUser, filter)
	if err == nil && filteredByUsernameUser.Username == user.Username {
		errorsFields["username"] = "Username has already been taken."
	} else if err != nil {
		log.Println(err)
		return err
	}

	if len(errorsFields) != 0 {
		return &data.BadInput{
			JSONError: berrors.BadInputsJSON(errorsFields),
			Err:       errors.New("register duplication error"),
		}
	}
	return nil
}
