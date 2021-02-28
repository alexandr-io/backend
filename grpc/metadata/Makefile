build:
	protoc --go_out=. --go_opt=module=github.com/alexandr-io/backend/grpc/grpcmetadata \
			--go-grpc_out=. --go-grpc_opt=module=github.com/alexandr-io/backend/grpc/grpcmetadata \
			protocol/metadata.proto
	goimports -w .
	python -m grpc_tools.protoc -I./protocol --python_out=. --grpc_python_out=. ./metadata/protocol/metadata.proto

.PHONY: build
