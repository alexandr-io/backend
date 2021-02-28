build:
	protoc --go_out=. --go_opt=module=github.com/alexandr-io/backend/grpc/grpcuser \
			--go-grpc_out=. --go-grpc_opt=module=github.com/alexandr-io/backend/grpc/grpcuser \
			protocol/user.proto
	goimports -w .

.PHONY: build
