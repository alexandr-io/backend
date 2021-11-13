package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/feedback"

	"github.com/gofiber/fiber/v2"
)

type field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	Color       int32     `json:"color"`
	Fields      []field   `json:"fields"`
}

type webhook struct {
	Content string  `json:"content"`
	Embeds  []embed `json:"embeds"`
}

// sendToDiscord sends a webhook request to discord to log user feedback
func sendToDiscord(userFeedback data.Feedback) error {
	body, err := json.Marshal(&webhook{
		Content: "Feedback received!",
		Embeds: []embed{
			{
				Title:       userFeedback.Title,
				Description: userFeedback.Content,
				Timestamp:   userFeedback.Timestamp,
				Color:       0xffca00,
				Fields: []field{
					{
						Name:   "Author",
						Value:  "`" + userFeedback.AuthorEmail + "`",
						Inline: true,
					},
					{
						Name:   "Device Info",
						Value:  "`" + userFeedback.AuthorDeviceInfo + "`",
						Inline: true,
					},
				},
			},
		},
	})
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	response, err := http.Post(os.Getenv("DISCORD_FEEDBACK_WEBHOOK"), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	log.Println(response.Status)
	log.Println(response)

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	log.Println(string(respBody))

	return nil
}

// UserFeedback sends a user feedback to the database and logs it in Discord
func UserFeedback(ctx *fiber.Ctx) error {
	var userFeedback data.Feedback
	if err := ParseBodyJSON(ctx, &userFeedback); err != nil {
		return err
	}

	userFeedback.Timestamp = time.Now()

	if userFeedback.Anonymous {
		userFeedback.AuthorEmail = "anonymous"
		userFeedback.AuthorDeviceInfo = "unknown"
	} else {
		if userFeedback.AuthorEmail == "" {
			return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "author_email must be provided if not anonymous")
		}
		if userFeedback.AuthorDeviceInfo == "" {
			return data.NewHTTPErrorInfo(fiber.StatusBadRequest, "author_device_info must be provided if not anonymous")
		}
	}

	insertedFeedback, err := feedback.Insert(userFeedback)
	if err != nil {
		return err
	}

	err = sendToDiscord(*insertedFeedback)
	if err != nil {
		return err
	}

	// Return the feedback data to the user
	if err := ctx.Status(fiber.StatusOK).JSON(insertedFeedback); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
