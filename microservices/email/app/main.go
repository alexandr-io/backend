// license that can be found in the LICENSE file.

// Mail is the alexandrio microservice that handle all the mails related features.
//
package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/alexandr-io/backend/mail/internal"
	"github.com/alexandr-io/backend/mail/kafka/consumers"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Mail Service started")
	log.Println("Mail Service started") // TODO: tmp

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sendgrid.DefaultClient = &rest.Client{HTTPClient: &http.Client{Transport: tr}}
	// Configure hermes by setting a theme and your product info
	internal.HMS = hermes.Hermes{
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name:      "Alexandrio",
			Link:      "http://alexandrio.cloud",
			Copyright: "Copyright © " + strconv.Itoa(time.Now().Year()) + " Alexandrio. All rights reserved.",
		},
	}
	for consumers.CreateTopics() != nil {
	}
	consumers.ConsumeEmailRequestMessages()
}
