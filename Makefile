include .env
export

export PROJECT_ROOT=$(shell pwd)

setup-hooks:
	git config core.hooksPath .githooks && \
	echo "Хуки настроены!"


cleanup-dp:
	@./.shell/cleanup-docker-postgres


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутвует переменная seq"; \
		echo "Example: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm migrate create -ext sql -dir /migrations -seq "$(seq)"


migrate-action:
	@if [ -z "$(act)" ]; then \
		echo "Отсутвует переменная act"; \
		echo "Example: make migrate-action act="[up, down] 1""; \
		exit 1; \
	fi; \
	docker compose run --rm migrate \
	-path /migrations \
	-database postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@postgres:5432/${DATABASE_NAME}?sslmode=disable \
	$(act)

migrate-up:
	@make migrate-action act=up

migrate-down:
	@make migrate-action act=down


run:
	@export LOGGER_PATH=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run ./cmd/app/main.go

test:
	@echo ${PROJECT_ROOT}