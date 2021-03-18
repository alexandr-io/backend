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

// UpdateEmailVerify create an email with the data.Email
func UpdateEmailVerify(mailData data.Email) error {
	// Create email object sender and receiver
	from := mail.NewEmail(os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_EMAIL"))
	subject := "Update your email"
	to := mail.NewEmail(mailData.Username, mailData.Email)

	// Create email content
	htmlContent, plainTextContent, err := createUpdateEmailVerifyBody(mailData)
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

// createUpdateEmailVerifyBody create the html and text body of the email
func createUpdateEmailVerifyBody(mailData data.Email) (string, string, error) {
	email := hermes.Email{
		Body: hermes.Body{
			Name: mailData.Username,
			Intros: []string{
				"You have requested to change your Alexandrio's email to this address",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To confirm this request, please use the following link (valid for 3 days):",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your email update",
						Link:  mailData.Data,
					},
				},
			},
			Outros: []string{
				"If you haven't made this request, or do not want to change your email address anymore, just ignore this e-mail.",
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
