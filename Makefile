APP_NAME=matcha

.PHONY: gqlgen bobgen network migrate-create migrate-up migrate-down

gqlgen:
	@go run github.com/99designs/gqlgen

bobgen:
	@PSQL_DSN=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/$(APP_NAME)?sslmode=disable go run github.com/stephenafamo/bob/gen/bobgen-psql@latest

network:
	@docker network create matcha-network

migrate-create:
	@migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	@migrate -path db/migrations -database 'pgx5://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/$(APP_NAME)' -verbose up

migrate-down:
	@migrate -path db/migrations -database 'pgx5://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/$(APP_NAME)' -verbose down
