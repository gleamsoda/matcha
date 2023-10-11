.PHONY: gqlgen

gqlgen:
	@gqlgen

dev-client:
	@pnpm -C ./web dev
