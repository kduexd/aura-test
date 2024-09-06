.PHONY: gen-swagger
gen-swagger:
	@swag init -g main.go -parseInternal --parseDependency --parseDepth 1 -o api/docs

.PHONY: migrate-create
migrate-create:
	@migrate create -ext sql -dir schema/migrations $(NAME)
	@echo "Migrate create" $(name) "successfully"

.PHONY: migrate-up
migrate-up:
	@migrate -database $(DB_URL) -path schema/migrations up
	@echo "Migrate up successfully"

.PHONY: migrate-down
migrate-down:
	@migrate -database $(DB_URL) -path schema/migrations down
	@echo "Migrate down successfully"

.PHONY: migrate-force
migrate-force:
	@migrate -database $(DB_URL) -path schema/migrations force $(VERSION)
	@echo "Migrate force successfully"