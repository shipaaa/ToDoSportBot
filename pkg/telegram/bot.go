package telegram

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	api         *tgbotapi.BotAPI
	dataBase    *sql.DB
	gendersUser map[string]string
}

func NewBot(bot *tgbotapi.BotAPI, dataBase *sql.DB) *Bot {
	return &Bot{api: bot, dataBase: dataBase}
}

func (b *Bot) Start() error {
	log.Infof("Authorized on account %s", b.api.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(-1) // Receive the last message from the user when we turn on the bot
	u.Timeout = 60
	return b.api.GetUpdatesChan(u)
}
