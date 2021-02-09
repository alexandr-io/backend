module github.com/alexandr-io/backend/library

go 1.15

replace github.com/alexandr-io/backend/library/data => ./data

require (
	github.com/alexandr-io/backend/library/data v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/berrors v1.2.7
	github.com/confluentinc/confluent-kafka-go v1.5.2
	github.com/fatih/structtag v1.2.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gofiber/fiber/v2 v2.3.3
	github.com/google/uuid v1.1.1
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	go.mongodb.org/mongo-driver v1.4.5
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
