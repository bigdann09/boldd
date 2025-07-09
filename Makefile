.PHONY: devrun migrate

devrun:
	go run cmd/api/main.go

migrate_up:
	migrate -path $(shell pwd)/internal/infrastructure/persistence/migrations -database "postgres://postgres:postgres@localhost:5432/boldd?sslmode=disable" up

migrate_down:
	# migrate -path $(shell pwd)/internal/infrastructure/persistence/migrations -database "postgres://postgres:postgres@localhost:5432/boldd?sslmode=disable" down

docker-migrate:
	docker run --rm -v $(shell pwd)/infrastructure/persistence/migrations:/migrations migrate/migrate -path=/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" up

docker-down:
	docker run --rm -v $(shell pwd)/infrastructure/persistence/migrations:/migrations migrate/migrate -path=/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" down

swag-generate:
	swag init -g cmd/api/main.go -o internal/api/docs