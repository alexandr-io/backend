mock:
	mkdir -p ./auth/mock ./library/mock ./user/mock
	mockgen -source=./auth/auth_grpc.pb.go AuthClient > auth/mock/auth_mock.go
	mockgen -source=./library/library_grpc.pb.go LibraryClient > library/mock/library_mock.go
	mockgen -source=./user/user_grpc.pb.go UserCLient > user/mock/user_mock.go

.PHONY: mock
