module github.com/alexandr-io/backend/auth/tests

go 1.15

replace github.com/alexandr-io/backend/tests/integrationMethods => ./../../../tests/integration_methods

replace github.com/alexandr-io/backend/auth/data => ../app/data

require (
	github.com/alexandr-io/backend/auth/data v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/tests/integrationMethods v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/berrors v1.2.7 // indirect
	github.com/confluentinc/confluent-kafka-go v1.5.2 // indirect
	github.com/gofiber/fiber/v2 v2.3.3 // indirect
	go.mongodb.org/mongo-driver v1.4.4 // indirect
)
