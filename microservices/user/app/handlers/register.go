package handlers

import (
	"github.com/gofiber/fiber"
)

// swagger:route POST /register USER register
// Register a new user and return it's information, auth token and refresh token
// responses:
//	201: userResponse
//	400: badRequestErrorResponse

// Register take a data.UserRegister in the body to create a new user in the database.
// The register route return a data.User.
func Register(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	//// Get and validate the body JSON
	//userRegister := new(data.UserRegister)
	//if ok := berrors.ParseBodyJSON(ctx, userRegister); !ok {
	//	return
	//}
	//
	//if userRegister.Password != userRegister.ConfirmPassword {
	//	errorData := berrors.BadInputJSON("confirm_password", "passwords does not match")
	//	ctx.Status(http.StatusBadRequest).SendBytes(errorData)
	//	return
	//}
	//
	//// Insert the new data to the collection
	//insertedResult := database.InsertUserRegister(ctx, data.User{
	//	Username: userRegister.Username,
	//	Email:    userRegister.Email,
	//	Password: hashAndSalt(userRegister.Password),
	//})
	//if insertedResult == nil {
	//	return
	//}
	//
	//// Get the newly created user
	//createdUser, ok := database.GetUserByID(ctx, insertedResult.InsertedID)
	//if !ok {
	//	return
	//}
	//
	//// Create auth and refresh token
	//refreshToken, authToken, ok := generateNewRefreshTokenAndAuthToken(ctx, createdUser.ID)
	//if !ok {
	//	return
	//}
	//createdUser.AuthToken = authToken
	//createdUser.RefreshToken = refreshToken
	//
	//// Return the new user to the user
	//if err := ctx.Status(201).JSON(createdUser); err != nil {
	//	berrors.InternalServerError(ctx, err)
	//}
	ctx.SendStatus(200)
}
