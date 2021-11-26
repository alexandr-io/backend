package internal

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alexandr-io/backend/mail/data"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// UpdatedPasswordMail create an email with the data.Email
func UpdatedPasswordMail(mailData data.Email) error {
	// Create email object sender and receiver
	from := mail.NewEmail(os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_EMAIL"))
	subject := "Your password has been modified"
	to := mail.NewEmail(mailData.Username, mailData.Email)
	// Create email content
	htmlContent, plainTextContent, err := createUpdatedPasswordBody(mailData)
	if err != nil {
		return err
	}
	// Send email
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	// Don't send email in test cases
	if !strings.Contains(mailData.Email, "@test.test") {
		resp, err := client.Send(message)
		if err != nil || resp.StatusCode != http.StatusAccepted {
			log.Printf("%v: %+v\n", err, resp)
			return err
		}
	}
	log.Printf("[MAIL] email sent to %s for %s", mailData.Email, mailData.Type)
	return nil
}

// createUpdatedPasswordBody create the html and text body of the email
func createUpdatedPasswordBody(mailData data.Email) (string, string, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Name: mailData.Username,
			Intros: []string{
				"Your password has been modified.",
			},
			Outros: []string{
				"If you haven't made this update, you can still change your password using the 'forgot my password' feature.",
			},
		},
	}

	// Generate an HTML email with the provided contents
	emailBody, err := HMS.GenerateHTML(email)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	// Generate the plaintext version of the e-mail
	emailText, err := HMS.GeneratePlainText(email)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	return emailBody, emailText, nil
}
