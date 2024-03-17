all: run

up:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down
run:
	go run application/*.go