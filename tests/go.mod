module github.com/alexandr-io/backend/tests

go 1.15

replace github.com/alexandr-io/backend/auth/tests => ../microservices/auth/tests

replace github.com/alexandr-io/backend/user/tests => ../microservices/user/tests

replace github.com/alexandr-io/backend/library/tests => ../microservices/library/tests

require (
	github.com/alexandr-io/backend/auth/tests v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/library/tests v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/user/tests v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.3.0
)
