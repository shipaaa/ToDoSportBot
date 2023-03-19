package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"strings"
)

func (b *Bot) sendMessage(chatID int64, messageText string) error {
	msg := tgbotapi.MessageConfig{BaseChat: tgbotapi.BaseChat{ChatID: chatID},
		ParseMode: "HTML", DisableWebPagePreview: true, Text: messageText}
	if _, err := b.api.Send(msg); err != nil {
		return err
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
	b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(b.dataBase,
		models.GenerateQuery(exercise))))
	return nil
}

func (b *Bot) sendMessageForKeyboardTraining(chatId int64, exercise string) error {
	exercise = strings.TrimSuffix(exercise, "Tr1")
	switch exercise {
	case "breast":
		b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(
			b.dataBase, models.GenerateQuery(exercise)+" LIMIT 4")))
	case "biceps":
		b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(
			b.dataBase, models.GenerateQuery(exercise)+" LIMIT 4")))
	case "back":
		b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(
			b.dataBase, models.GenerateQuery(exercise)+" LIMIT 5")))
	case "triceps":
		b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(
			b.dataBase, models.GenerateQuery(exercise)+" LIMIT 3")))
	case "leg":
		b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(
			b.dataBase, models.GenerateQuery(exercise)+" LIMIT 5")))
	case "shoulder":
		b.sendMessage(chatId, getMessageTextForExerciseKeyboard(models.GetDataFromDB(
			b.dataBase, models.GenerateQuery(exercise)+" LIMIT 3")))
	}
	return nil
}

func (b *Bot) sendWaitingMessage(keyboard *tgbotapi.InlineKeyboardMarkup,
	callbackQuery *tgbotapi.CallbackQuery) string {
	exerciseMessage := fmt.Sprintf(messageAboutSendEx, getExerciseName(keyboard, callbackQuery))
	b.sendMessage(callbackQuery.Message.Chat.ID, exerciseMessage)
	return callbackQuery.Data
}
