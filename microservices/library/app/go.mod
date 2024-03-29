module github.com/alexandr-io/backend/library

go 1.15

replace github.com/alexandr-io/backend/grpc => ../../../grpc

replace github.com/alexandr-io/backend/library/data => ./data

replace github.com/alexandr-io/backend/common => ../../../common

require (
	github.com/alexandr-io/backend/common v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/grpc v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/library/data v0.0.0-00010101000000-000000000000
	github.com/fatih/structtag v1.2.0
	github.com/getsentry/sentry-go v0.11.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.0
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.5.2
	google.golang.org/grpc v1.35.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
