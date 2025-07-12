DB_URL=postgres://postgres:postgres@localhost:5432/gofinance?sslmode=disable
MIGRATE_PATH=server/db/migration

.PHONY: migrate-up migrate-down migrate-create migrate-status run-server run-client protoc

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" down

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Please provide migration name: make migrate-create name=init"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(name)

migrate-status:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" version

run-server:
	go run ./server/cmd/main.go

run-client:
	go run ./client/cmd/main.go

protoc:
	protoc --go_out=. --go-grpc_out=. --proto_path=. shared/proto/transaction.proto