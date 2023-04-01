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
	query := models.GenerateQuery(exercise)
	tableData, err := models.GetDataFromDB(b.dataBase, query)
	if err != nil {
		return err
	}
	b.sendMessage(chatId, getMessageTextForExerciseKeyboard(tableData))
	return nil
}

func (b *Bot) sendMessageForKeyboardTraining(chatId int64, exercise string) error {
	exercise = strings.TrimSuffix(exercise, "Tr1")
	exerciseLimit := map[string]int{
		"breast":   4,
		"biceps":   4,
		"back":     5,
		"triceps":  3,
		"leg":      5,
		"shoulder": 3,
	}
	limit, ok := exerciseLimit[exercise]
	if !ok {
		return fmt.Errorf("unsupported exercise type: %s", exercise)
	}
	query := models.GenerateQuery(exercise) + fmt.Sprintf(" LIMIT %d", limit)
	data, err := models.GetDataFromDB(b.dataBase, query)
	if err != nil {
		return err
	}
	text := getMessageTextForExerciseKeyboard(data)
	b.sendMessage(chatId, text)
	return nil
}

func (b *Bot) sendWaitingMessage(keyboard *tgbotapi.InlineKeyboardMarkup,
	callbackQuery *tgbotapi.CallbackQuery) string {
	exerciseMessage := fmt.Sprintf(msgAboutSendEx, getExerciseName(keyboard, callbackQuery))
	b.sendMessage(callbackQuery.Message.Chat.ID, exerciseMessage)
	return callbackQuery.Data
}
