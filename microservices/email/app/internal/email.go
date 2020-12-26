package internal

import (
	"log"

	"github.com/alexandr-io/backend/mail/data"

	"github.com/matcornic/hermes/v2"
)

// HMS is a global variable used to create email with hermes. It is configured in the main function
var HMS hermes.Hermes

type mailFunc func(data.KafkaEmail) error

// messageTypeMap map of mail types and their corresponding function
var messageTypeMap = map[string]mailFunc{
	"password-reset": ResetPasswordMail,
}

// CreateMailFromMessage is reading the data.KafkaEmail.Type to execute the corresponding function stored in messageTypeMap
func CreateMailFromMessage(mail data.KafkaEmail) {
	function, ok := messageTypeMap[mail.Type]
	if !ok {
		log.Printf("Mail type %s not recognized\n", mail.Type)
		return
	}

	_ = function(mail)
}
