run-grpc:
	cd service-core && go run ./cmd/grpc/main.go

test-service-core:
	cd service-core && go test ./...

gen-grpc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto	