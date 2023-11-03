APP_NAME=matcha

.PHONY: gqlgen network migrate-create migrate-up migrate-down

gqlgen:
	@gqlgen

network:
	@docker network create matcha-network

migrate-create:
	@migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	@migrate -path db/migrations -database 'pgx5://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/$(APP_NAME)' -verbose up

migrate-down:
	@migrate -path db/migrations -database 'pgx5://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/$(APP_NAME)' -verbose down
