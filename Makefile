.PHONY: up down producer consumer test
up:
	docker compose up -d
down:
	docker compose down
producer:
	go run ./cmd/producer
consumer:
	go run ./cmd/consumer
test:
	go test ./...
