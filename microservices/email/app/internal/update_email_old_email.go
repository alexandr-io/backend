package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alexandr-io/backend/mail/data"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// UpdateEmailOldEmail create an email with the data.Email
func UpdateEmailOldEmail(mailData data.Email) error {
	// Create email object sender and receiver
	from := mail.NewEmail(os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_EMAIL"))
	subject := "EMAIL CHANGED"
	to := mail.NewEmail(mailData.Username, mailData.Email)

	// Create email content
	htmlContent, plainTextContent, err := createUpdateEmailOldEmailBody(mailData)
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

// createUpdateEmailOldEmailBody create the html and text body of the email
func createUpdateEmailOldEmailBody(mailData data.Email) (string, string, error) {
	var emailData struct {
		NewEmail string `json:"new_email"`
		Link     string `json:"link"`
	}
	if err := json.Unmarshal([]byte(mailData.Data), &emailData); err != nil {
		return "", "", err
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: mailData.Username,
			Intros: []string{
				"A request have been made on your account to change the email address to " + emailData.NewEmail,
			},
			Actions: []hermes.Action{
				{
					Instructions: "You can cancel this update by using the following link (valid for 3 days):",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Cancel the email update",
						Link:  emailData.Link,
					},
				},
			},
			Outros: []string{
				"If you haven't made this request, or if you do want to change your email address, just ignore this e-mail.",
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
