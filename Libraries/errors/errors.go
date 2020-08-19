package errors

import (
	"encoding/json"
	"log"
)

func BadInputsJSON(fields map[string]string) string {
	badInputData := badInput{}
	for key, element := range fields {
		badInputData.Fields = append(badInputData.Fields, field{
			Name:   key,
			Reason: element,
		})
	}

	log.Println(json.Marshal(badInputData))
	return ""
}
