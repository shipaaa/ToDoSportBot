package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"github.com/shipaaa/telegram-sport-bot/pkg/telegram"

	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	db, err := models.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bot := telegram.NewBot(botAPI, db)
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
