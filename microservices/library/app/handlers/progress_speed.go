package handlers

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	progressSpeedServ "github.com/alexandr-io/backend/library/internal/progressspeed"
	userMiddleware "github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
)

func progressMade(ctx *fiber.Ctx) error {
	// Parse data
	var newProgress data.NewProgress
	if err := ParseBodyJSON(ctx, &newProgress); err != nil {
		return err
	}

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	if err = progressSpeedServ.Serv.UpsertProgressSpeed(userID, newProgress.Language, newProgress.WordNumber); err != nil {
		return err
	}

	if err = ctx.SendStatus(fiber.StatusCreated); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func readingSpeed(ctx *fiber.Ctx) error {
	// Parse data
	var getReadingSpeed data.GetReadingSpeed
	if err := ParseBodyJSON(ctx, &getReadingSpeed); err != nil {
		return err
	}

	// Get data from header and params
	userID, err := userIDFromHeader(ctx)
	if err != nil {
		return err
	}

	readingSpeed, err := progressSpeedServ.Serv.ReadReadingSpeed(userID, getReadingSpeed.Language, getReadingSpeed.WordNumber)
	if err != nil {
		return err
	}

	if err = ctx.Status(fiber.StatusOK).JSON(readingSpeed); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// CreateProgressSpeedHandlers create handlers for ProgressSpeed
func CreateProgressSpeedHandlers(app *fiber.App) {
	progressSpeedServ.NewService(database.ProgressSpeedDB)
	app.Post("/progress/speed", userMiddleware.Protected(), progressMade)
	app.Get("/progress/speed", userMiddleware.Protected(), readingSpeed)
}
