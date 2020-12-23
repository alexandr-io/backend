package internal

import (
	"log"

	"github.com/alexandr-io/backend/mail/data"

	"github.com/matcornic/hermes/v2"
)

var HMS hermes.Hermes

type mailFunc func(data.KafkaMail) error

var messageTypeMap = map[string]mailFunc{
	"password-reset": ResetPasswordMail,
}

func CreateMailFromMessage(mail data.KafkaMail) {
	function, ok := messageTypeMap[mail.Type]
	if !ok {
		log.Printf("Mail type %s not recognized\n", mail.Type)
		return
	}

	_ = function(mail)
}
