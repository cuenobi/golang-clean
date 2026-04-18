APP_NAME=golang-clean

.PHONY: help run-api run-consumer migrate-up migrate-down test mockery fmt

help:
	@echo "make run-api        - run HTTP API"
	@echo "make run-consumer   - run Kafka consumer"
	@echo "make migrate-up     - apply DB migrations"
	@echo "make migrate-down   - rollback DB migration"
	@echo "make test           - run unit tests"
	@echo "make mockery        - generate mocks"
	@echo "make fmt            - format Go files"

run-api:
	go run ./cmd/app api

run-consumer:
	go run ./cmd/app consumer

migrate-up:
	go run ./cmd/app migrate up

migrate-down:
	go run ./cmd/app migrate down 1

test:
	go test ./...

mockery:
	go generate ./...

fmt:
	gofmt -w $(shell find . -type f -name '*.go' -not -path './vendor/*')
