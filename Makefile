all: run

up:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down

run:
	go run application/*.go

docker-build:
	docker build -t ta-bc-dashboard .
docker-run:
	docker run -d -p 8080:8080 -p 9090:9090 ta-bc-dashboard