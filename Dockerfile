FROM golang:1.20.2-alpine3.16 AS builder

COPY . /github.com/shipaaa/telegram-sport-bot
WORKDIR /github.com/shipaaa/telegram-sport-bot

RUN go mod tidy
RUN go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/shipaaa/telegram-sport-bot/.bin/bot .
COPY .env .

EXPOSE 80

CMD ["./bot"]
