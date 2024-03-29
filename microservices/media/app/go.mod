module github.com/alexandr-io/backend/media

go 1.15

replace github.com/alexandr-io/backend/grpc => ../../../grpc

require (
	github.com/alexandr-io/backend/grpc v0.0.0-00010101000000-000000000000
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/aws/aws-sdk-go v1.37.9 // indirect
	github.com/fatih/structtag v1.2.0
	github.com/getsentry/sentry-go v0.11.0
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/golang/mock v1.5.0
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/h2non/filetype v1.1.1
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.20.0 // indirect
	go.mongodb.org/mongo-driver v1.5.1
	go.opencensus.io v0.22.6 // indirect
	gocloud.dev v0.22.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	google.golang.org/genproto v0.0.0-20210211221406-4ccc9a5e4183 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/go-playground/validator.v9 v9.31.0
)
