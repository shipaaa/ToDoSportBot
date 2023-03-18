FROM golang:1.20.2-alpine3.17 AS builder

COPY . /github.com/shipaaa/telegram-sport-bot
WORKDIR /github.com/shipaaa/telegram-sport-bot

RUN go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/shipaaa/telegram-sport-bot/.bin/bot .