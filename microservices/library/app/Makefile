mock:
	mkdir -p ./internal/book/mock ./internal/bookprogress/mock ./internal/group/mock ./internal/library/mock ./internal/metadata/mock ./internal/permission/mock ./internal/userdata/mock ./internal/userlibrary/mock
	mockgen -source=./internal/group/interface.go > ./internal/group/mock/mock.go
	mockgen -source=./internal/userlibrary/interface.go > ./internal/userlibrary/mock/mock.go

.PHONY: mock
