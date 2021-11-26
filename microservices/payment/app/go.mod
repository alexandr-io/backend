module github.com/alexandr-io/backend/payment

go 1.16

replace github.com/alexandr-io/backend/grpc => ../../../grpc

replace github.com/alexandr-io/backend/common => ../../../common

require (
	github.com/alexandr-io/backend/common v0.0.0-20210913201527-0adcfe9625ce
	github.com/alexandr-io/backend/grpc v0.0.0-00010101000000-000000000000
	github.com/fatih/structtag v1.2.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gofiber/fiber/v2 v2.14.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/stripe/stripe-go/v72 v72.56.0
	go.mongodb.org/mongo-driver v1.5.4
	google.golang.org/grpc v1.39.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)
