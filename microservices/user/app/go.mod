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
	github.com/confluentinc/confluent-kafka-go v1.5.2
	github.com/fatih/structtag v1.2.0
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/golang/snappy v0.0.2 // indirect
	github.com/klauspost/compress v1.11.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	go.mongodb.org/mongo-driver v1.4.3
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.35.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
