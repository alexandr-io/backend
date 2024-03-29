module github.com/alexandr-io/backend/auth

go 1.15

replace github.com/alexandr-io/backend/auth/data => ./data

replace github.com/alexandr-io/backend/grpc => ../../../grpc

replace github.com/alexandr-io/backend/common => ../../../common

require (
	github.com/alexandr-io/backend/auth/data v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/common v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/grpc v0.0.0-00010101000000-000000000000
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/fatih/structtag v1.2.0
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible
	github.com/getsentry/sentry-go v0.11.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-redis/redismock/v8 v8.0.6
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/gofiber/jwt/v2 v2.1.0
	github.com/golang/mock v1.5.0
	github.com/klauspost/compress v1.11.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.18.0
	go.mongodb.org/mongo-driver v1.5.2
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
