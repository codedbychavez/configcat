docker:
	docker-compose up --build

env:
	cp .env.example .env

install:
	go mod download

run:
	go run .