module github.com/alexandr-io/backend/tests/integration

go 1.15

replace github.com/alexandr-io/backend/tests/itgmtod => ./../itgmtod

replace github.com/alexandr-io/backend/auth/tests => ../../microservices/auth/tests

replace github.com/alexandr-io/backend/user/tests => ../../microservices/user/tests

replace github.com/alexandr-io/backend/library/tests => ../../microservices/library/tests

replace github.com/alexandr-io/backend/auth/data => ../../microservices/auth/app/data

replace github.com/alexandr-io/backend/library/data => ../../microservices/library/app/data

replace github.com/alexandr-io/backend/common => ../../common

require (
	github.com/alexandr-io/backend/auth/tests v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/common v0.0.0-00010101000000-000000000000 // indirect
	github.com/alexandr-io/backend/library/tests v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/user/tests v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.3.0
)
