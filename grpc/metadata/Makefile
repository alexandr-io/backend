build:
	protoc --go_out=. --go_opt=module=github.com/alexandr-io/backend/grpc/grpcauth \
			--go-grpc_out=. --go-grpc_opt=module=github.com/alexandr-io/backend/grpc/grpcauth \
			./metadata/protocol/metadata.proto
	goimports -w .
	python -m grpc_tools.protoc -I../../protos --python_out=. --grpc_python_out=. ./metadata/protocol/metadata.proto

.PHONY: build
