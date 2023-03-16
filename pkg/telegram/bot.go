package telegram

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	dataBase *sql.DB
}

func NewBot(bot *tgbotapi.BotAPI, dataBase *sql.DB) *Bot {
	return &Bot{bot: bot, dataBase: dataBase}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) send(chatID int64, message string) error {
	msg := tgbotapi.MessageConfig{BaseChat: tgbotapi.BaseChat{ChatID: chatID},
		ParseMode: "HTML", DisableWebPagePreview: true, Text: message}
	if _, err := b.bot.Send(msg); err != nil {
		b.bot.Send(tgbotapi.NewMessage(chatID, "Произошла ошибка, уже исправляем"))
		return err
	}
	return nil
}
