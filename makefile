.PHONY: psql
psql:
	docker exec -it service-core-db psql -U postgres plataforma-ead

.PHONY: up
up:
	if [ "$(MODE)" = "l" ]; then \
		docker compose up; \
	else \
		docker compose up -d; \
	fi

.PHONY: restart
restart:
	docker compose restart

.PHONY: log
log:
	docker logs -f service-core

.PHONY: down

down:
	docker compose down


run-grpc:
	cd service-core && go run ./cmd/grpc/main.go

test-service-core:
	cd service-core && go test ./...

gen-grpc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto
   
