.PHONY: tests

COMPOSE_FILE=docker-compose.test.yaml
PROJECT_NAME=hiven-tests

tests:
	@echo Creating environment...
	@docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d
	@echo Starting tests...
	@docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) logs -f api-server
	@echo Cleaning up...
	@docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down -v
