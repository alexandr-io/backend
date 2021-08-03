package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	
	"github.com/alexandr-io/backend/library/data"
	"github.com/gofiber/fiber/v2"
)

// DictionaryRetrieve gets the definition of a queried word with a queried language
func DictionaryRetrieve(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	queriedWord := ctx.Params("queried_word")
	queriedLang := ctx.Params("lang")

	url := "https://api.dictionaryapi.dev/api/v2/entries/" + queriedLang + "/" + queriedWord

	response, err := http.Get(url)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	definitionByes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	var dictResponse []data.DictResponse

	if response.Body != nil && response.StatusCode == 200 {
		err := json.Unmarshal(definitionByes, &dictResponse)
		if err != nil {
			return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
	}

	if err = ctx.Status(fiber.StatusOK).JSON(dictResponse); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
