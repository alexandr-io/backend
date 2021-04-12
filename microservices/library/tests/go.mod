module github.com/alexandr-io/backend/library/tests

go 1.15

replace github.com/alexandr-io/backend/tests/itgmtod => ./../../../tests/itgmtod

replace github.com/alexandr-io/backend/library/data => ../app/data

replace github.com/alexandr-io/backend/common => ../../../common

require (
	github.com/alexandr-io/backend/common v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/library/data v0.0.0-00010101000000-000000000000
	github.com/alexandr-io/backend/tests/itgmtod v0.0.0-00010101000000-000000000000
)
