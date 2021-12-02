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

	"github.com/alexandr-io/backend/mail/grpc"
	"github.com/alexandr-io/backend/mail/internal"

	"github.com/getsentry/sentry-go"
	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Mail Service started")

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://a3a4b09e28514398a147c7ac0c43eedf@95.217.135.159:9000/2",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sendgrid.DefaultClient = &rest.Client{HTTPClient: &http.Client{Transport: tr}}
	// Configure hermes by setting a theme and your product info
	internal.HMS = hermes.Hermes{
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name:      "Alexandrio",
			Link:      "https://alexandrio.cloud",
			Copyright: "Copyright © " + strconv.Itoa(time.Now().Year()) + " Alexandrio. All rights reserved.",
		},
	}

	grpc.InitGRPC()
}
