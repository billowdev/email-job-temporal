
docker:
	@docker compose up

run:
	@echo "Running the application..."
	@go run cmd/main.go

.PHONY: run docker