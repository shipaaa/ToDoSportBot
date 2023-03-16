package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"time"
)

func (b *Bot) sendMessage(chatID int64, messageText string) error {
	msg := tgbotapi.MessageConfig{BaseChat: tgbotapi.BaseChat{ChatID: chatID},
		ParseMode: "HTML", DisableWebPagePreview: true, Text: messageText}
	if _, err := b.api.Send(msg); err != nil {
		return b.sendMessage(chatID, "Произошла ошибка, уже исправляем")
	}
	return nil
}

func (b *Bot) deleteMessage(chatID int64, messageID int) error {
	del := tgbotapi.NewDeleteMessage(chatID, messageID)
	if _, err := b.api.Request(del); err != nil {
		return err
	}
	return nil
}

func (b *Bot) sendKeyboard(chatID int64, text string, keyboardFunc func() tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboardFunc()
	if _, err := b.api.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) sendMessageForExerciseKeyboard(chatId int64, exercise string) error {
	time.Sleep(1 * time.Second)
	b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetAllExercisesFromDB(b.dataBase, exercise)))
	return nil
}

func (b *Bot) sendWaitingMessage(keyboard *tgbotapi.InlineKeyboardMarkup,
	callbackQuery *tgbotapi.CallbackQuery) (int64, string) {
	exerciseMessage := fmt.Sprintf("Присылаю упражнения на <b>%s</b>\n"+
		"Секундочку...", getExerciseName(keyboard, callbackQuery))
	b.sendMessage(callbackQuery.Message.Chat.ID, exerciseMessage)
	return callbackQuery.Message.Chat.ID, callbackQuery.Data
}
