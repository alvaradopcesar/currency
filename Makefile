
init:
	cp .env.template .env

init-compose:
	docker compose -f docker-compose.yml up -d;

stop-compose:
	docker compose stop

start-compose:
	docker-compose up -d

down-compose:
	docker compose down

start:
	go run cmd/api/main.go

test:
	@go test -cover ./...