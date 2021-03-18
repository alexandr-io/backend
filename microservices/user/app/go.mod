module github.com/alexandr-io/backend/user

go 1.15

replace github.com/alexandr-io/backend/grpc => ../../../grpc

replace github.com/alexandr-io/backend/common => ../../../common

require (
	github.com/alexandr-io/backend/common v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/grpc v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/berrors v1.2.7
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/aws/aws-sdk-go v1.35.35 // indirect
	github.com/fatih/structtag v1.2.0
	github.com/go-redis/redis/v8 v8.7.1
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/gofiber/template v1.6.7
	github.com/golang/snappy v0.0.2 // indirect
	github.com/klauspost/compress v1.11.3 // indirect
	go.mongodb.org/mongo-driver v1.4.3
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/grpc v1.35.0
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
