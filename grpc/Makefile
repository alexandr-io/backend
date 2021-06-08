mock:
	mkdir -p ./auth/mock ./library/mock ./user/mock ./email/mock ./metadata/mock
	mockgen -source=./auth/auth_grpc.pb.go AuthClient > auth/mock/auth_mock.go
	mockgen -source=./library/library_grpc.pb.go LibraryClient > library/mock/library_mock.go
	mockgen -source=./user/user_grpc.pb.go UserCLient > user/mock/user_mock.go
	mockgen -source=./email/email_grpc.pb.go EmailClient > email/mock/email_mock.go
	mockgen -source=./metadata/metadata_grpc.pb.go MetadataClient > metadata/mock/metadata_mock.go

.PHONY: mock
