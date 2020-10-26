package handlers

import (
	"encoding/json"
	"errors"
	"github.com/alexandr-io/backend/library/data"
	"github.com/fatih/structtag"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"reflect"
	"strings"
)

// getJSONFieldName is used to get the json tag of a given field in a struct
func getJSONFieldName(object interface{}, fieldName string) (string, error) {
	field, ok := reflect.TypeOf(object).Elem().FieldByName(fieldName)
	if !ok {
		return "", errors.New("Field name `" + fieldName + "` not found within the given object `" + reflect.TypeOf(object).Elem().Name() + "`")
	}
	tags, err := structtag.Parse(string(field.Tag))
	if err != nil {
		log.Println(err)
		return "", err
	}
	jsonTag, err := tags.Get("json")
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Remove `,omitempty` from the tag
	jsonTagName := strings.Replace(jsonTag.Name, ",omitempty", "", -1)

	return jsonTagName, nil
}

// ParseBodyJSON parse and validate a body contained in the fiber context to the given object
// If an error occur, the correct http error is called and false is returned
// The validator errors messages are using BadInputsJSONFromType
func ParseBodyJSON(ctx *fiber.Ctx, object interface{}) error {
	if err := ctx.BodyParser(object); err != nil {
		errorInfo := data.NewErrorInfo(err.Error(), 0)
		errorInfo.CustomMessage = "Error while parsing the json"
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	}

	v := validator.New()
	if err := v.Struct(object); err != nil {
		log.Println(err)
		errorMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			jsonTagName, err := getJSONFieldName(object, e.Field())
			if err != nil {
				return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
			}
			errorMap[jsonTagName] = e.Tag()
		}
		errorInfo := data.NewErrorInfo(string(badInputsJSONFromType(errorMap)), 0)
		errorInfo.ContentType = fiber.MIMEApplicationJSON
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	}
	return nil
}

// badInputsJSON creates the error JSON using the struct BadInput.
// The key of the map given correspond to the Name and the value to the Reason.
// It returns the JSON in []byte.
func badInputsJSON(fields map[string]string) []byte {
	badInputData := data.BadInput{}
	for key, element := range fields {
		badInputData.Fields = append(badInputData.Fields, data.Field{
			Name:   key,
			Reason: element,
		})
	}

	jsonData, _ := json.Marshal(badInputData)
	return jsonData
}

// badInputJSON is simply a call to BadInputsJSON to create a single bad input error.
// It returns the JSON of the struct BadInput in []byte.
func badInputJSON(name string, reason string) []byte {
	return badInputsJSON(map[string]string{name: reason})
}

// badInputsJSONFromType create a BadInput JSON from a key and a value corresponding to an ErrorType.
// It replace the Value with the defined string corresponding to the ErrorType.
// It returns the JSON in []byte.
func badInputsJSONFromType(fields map[string]string) []byte {
	newFields := make(map[string]string)
	for key, element := range fields {
		newFields[key] = data.ErrorTypes[data.ErrorType(element)]
	}
	return badInputsJSON(newFields)
}
