include .env
DOCKER_COMMAND=docker-compose
MIGRATE=${DOCKER_COMMAND} exec web migrate -path=migrations -database "postgres://${DBUsername}:${DBPassword}@${DBHost}:${DBPort}/${DBName}?sslmode=disable" -verbose

migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create-migration:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir migrations  $$NAME

create-code:
	${DOCKER_COMMAND} exec web go run ./cmd/. create_code

test-all:
	${DOCKER_COMMAND} exec web go test ./tests/tests/...

test-all-debugger:
	${DOCKER_COMMAND} exec web dlv test ./tests/tests/ --headless --listen=:4000 --api-version=2 --accept-multiclient

kill-test-debugger:
	${DOCKER_COMMAND} exec web pkill -f "dlv test"

.PHONY: migrate-up migrate-down force goto drop create

.PHONY: migrate-up migrate-down force goto drop create auto-create
