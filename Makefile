BIN=$(PWD)/bin
BUILD = $(PWD)/build/exercise
MAIN = $(PWD)/cmd/exercise

MIGRATIONS=$(PWD)/migrations

include .env

bin/migrate:
	GOBIN=$(BIN) go install -tags='postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: build
build:
	go build -tags musl -mod=vendor -o ${BUILD} ${MAIN}

.PHONY: run
run: build
	${BUILD}

.PHONY: run-env
run-env:
	docker-compose up -d

.PHONY: migrate-new
migrate-new: bin/migrate
	$(BIN)/migrate create -ext sql -seq -dir $(MIGRATIONS) '$(name)'

# run migrations
.PHONY: migrate-up
migrate-up: bin/migrate
	$(BIN)/migrate -path=$(MIGRATIONS) -database "${MIGRATE_ENDPOINT}" up

# down migrations
.PHONY: migrate-down
migrate-down: bin/migrate
	$(BIN)/migrate -path=$(MIGRATIONS) -database "${MIGRATE_ENDPOINT}" down -all