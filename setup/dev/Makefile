.PHONY: start stop

COMPOSE_FILE=docker-compose.dev.yaml
PROJECT_NAME=hiven-dev

start:
	@echo Starting development environment...
	@echo The API server may take a few seconds to become available...
	@docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d

stop:
	@echo Stopping development environment...
	@docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down
