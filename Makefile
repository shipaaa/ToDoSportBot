.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t shipaaa/telegram-sport-bot:0.1 .

start-container:
	docker run --env-file .env -p 80:80 shipaaa/telegram-sport-bot:0.1
