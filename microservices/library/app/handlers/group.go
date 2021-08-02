package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"
	groupServ "github.com/alexandr-io/backend/library/internal/group"
	permissionServ "github.com/alexandr-io/backend/library/internal/permission"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

// groupCreate create a group
func groupCreate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanManagePermissions() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to create a group in this library")
	}

	var group permissions.Group
	if err = ParseBodyJSON(ctx, &group); err != nil {
		return err
	}
	group.LibraryID = libraryID

	groupResult, err := groupServ.Serv.CreateGroup(group)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusCreated).JSON(groupResult); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// groupRetrieve retrieve a group
func groupRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}
	groupID, err := getGroupIDFromParams(ctx)
	if err != nil {
		return err
	}

	group, err := groupServ.Serv.ReadGroupFromIDAndLibraryID(groupID, libraryID)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(group); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// groupUpdate update a group
func groupUpdate(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}
	groupID, err := getGroupIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanManagePermissions() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to add a user in a group in this library")
	}

	var group permissions.Group
	if err := ParseBodyJSON(ctx, &group); err != nil {
		return err
	}
	group.LibraryID = libraryID
	group.ID = groupID

	groupResult, err := groupServ.Serv.UpdateGroup(group)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(groupResult); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// groupDelete delete a group
func groupDelete(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}
	groupID, err := getGroupIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanManagePermissions() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to delete a group in this library")
	}

	err = groupServ.Serv.DeleteGroup(groupID)
	if err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// groupAddUser add a user to a group
func groupAddUser(ctx *fiber.Ctx) error {

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}
	groupID, err := getGroupIDFromParams(ctx)
	if err != nil {
		return err
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return err
	} else if !perm.CanManagePermissions() {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to add a user in a group in this library")
	}

	var groupJoinData data.GroupJoin
	if err = ParseBodyJSON(ctx, &groupJoinData); err != nil {
		return err
	}

	if err = groupServ.Serv.AddUserToGroup(userID, groupID, libraryID); err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

// groupsRetrieveUser retrieve the list of the user's groups
func groupsRetrieveUser(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}
	libraryID, err := getLibraryIDFromParams(ctx)
	if err != nil {
		return err
	}

	result, err := groupServ.Serv.ReadGroupFromLibrary(libraryID, userID)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func CreateGroupHandlers(app *fiber.App) {
	groupServ.NewService(database.GroupDB, database.UserLibraryDB)
	app.Post("/library/:library_id/group", userMiddleware.Protected(), groupCreate)
	app.Get("/library/:library_id/group/:group_id", userMiddleware.Protected(), groupRetrieve)
	app.Post("/library/:library_id/group/:group_id", userMiddleware.Protected(), groupUpdate)
	app.Delete("/library/:library_id/group/:group_id", userMiddleware.Protected(), groupDelete)

	app.Post("/library/:library_id/group/:group_id/join", userMiddleware.Protected(), groupAddUser)
	app.Get("/library/:library_id/user/groups", userMiddleware.Protected(), groupsRetrieveUser)
}
