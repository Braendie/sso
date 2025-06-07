APP_NAME=sso
PATH_TO_MAIN=./cmd/sso/main.go
PATH_TO_CONFIG=./config/local.yaml
PATH_TO_MIGRATOR=./cmd/migrator/main.go
PATH_TO_STORAGE=./storage/sso.db
PATH_TO_MIGRATIONS=./migrations
PATH_TO_MIGRATIONS_TEST=./tests/migrations

.PHONY: build run migrator migrator_test test clean

build:
	go build -o ./bin/$(APP_NAME) $(PATH_TO_MAIN) -config

run: 
	go run $(PATH_TO_MAIN) -config $(PATH_TO_CONFIG)

migrator:
	go run $(PATH_TO_MIGRATOR) --storage-path=$(PATH_TO_STORAGE) --migrations-path=$(PATH_TO_MIGRATIONS)

migrator_test:
	go run $(PATH_TO_MIGRATOR) --storage-path=./storage/sso.db --migrations-path=$(PATH_TO_MIGRATIONS_TEST) --migrations-table=migrations_test

test:
	go test ./...

clean:
	rm -rf bin/

help:
	@echo "Использование:"
	@echo "  make run       - Запуск приложения"
	@echo "  make build     - Сборка бинарного файла"
	@echo "  make clean     - Удаление собранных файлов"
	@echo "  make test      - Запуск тестов"