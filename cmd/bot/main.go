package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shipaaa/telegram-sport-bot/pkg/config"
	"github.com/shipaaa/telegram-sport-bot/pkg/logging"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"github.com/shipaaa/telegram-sport-bot/pkg/telegram"
	log "github.com/sirupsen/logrus"
)

func main() {
	logging.StartLogging()

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	db, err := models.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bot := telegram.NewBot(botAPI, db)
	if err = bot.Start(); err != nil {
		log.Fatal(err)
	}
}
