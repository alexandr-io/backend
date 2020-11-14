module github.com/alexandr-io/backend/tests

go 1.15

replace github.com/alexandr-io/backend/auth/tests => ../microservices/auth/tests

require (
	github.com/alexandr-io/backend/auth/tests v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.3.0
)
