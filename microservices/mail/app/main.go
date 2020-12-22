// license that can be found in the LICENSE file.

// Mail is the alexandrio microservice that handle all the mails related features.
//
package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Mail Service started")

	from := mail.NewEmail("Alexandrio", "no-reply@alexandrio.cloud")
	subject := "Reset password"
	to := mail.NewEmail("Alexandrio", "alexandriocloud@gmail.com")
	plainTextContent := "Reset your password by clicking this link"
	htmlContent := "Reset your password by clicking this link"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
