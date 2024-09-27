DB_DSN := "postgres://postgres:postgres@localhost:5432/2task?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

lint:
	golangci-lint run --out-format=colored-line-number

gen:
	oapi-codegen -config openapi/.openapi openapi/openapi.yaml > ./internal/api/api.gen.go

run:
	go run cmd/app/main.go 
