build:
	protoc --go_out=. --go_opt=module=github.com/alexandr-io/backend/grpc/grpcauth \
			--go-grpc_out=. --go-grpc_opt=module=github.com/alexandr-io/backend/grpc/grpcauth \
			protocol/auth.proto
	goimports -w .
