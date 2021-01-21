module github.com/alexandr-io/backend/auth/tests

go 1.15

replace github.com/alexandr-io/backend/tests/itgmtod => ./../../../tests/itgmtod

replace github.com/alexandr-io/backend/auth/data => ../app/data

require (
	github.com/alexandr-io/backend/auth/data v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/tests/itgmtod v0.0.0-00010101000000-000000000000
)
