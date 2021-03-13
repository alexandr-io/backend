build:
	protoc --go_out=. --go_opt=module=github.com/alexandr-io/backend/grpc/grpcemail \
			--go-grpc_out=. --go-grpc_opt=module=github.com/alexandr-io/backend/grpc/grpcemail \
			protocol/email.proto
	goimports -w .

.PHONY: build
