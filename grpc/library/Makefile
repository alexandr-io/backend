build:
	protoc --go_out=. --go_opt=module=github.com/alexandr-io/backend/grpc/grpclibrary \
			--go-grpc_out=. --go-grpc_opt=module=github.com/alexandr-io/backend/grpc/grpclibrary \
			protocol/library.proto
	goimports -w .

.PHONY: build
