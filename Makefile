.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

up:
	docker-compose -f docker-compose.yml up -d

start:
	docker-compose -f docker-compose.yml start

stop:
	docker-compose -f docker-compose.yml stop