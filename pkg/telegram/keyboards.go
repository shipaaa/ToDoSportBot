package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"time"
)

func (b *Bot) keyboardAllExercises() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Грудь", "breast"),
			tgbotapi.NewInlineKeyboardButtonData("Бицепс", "biceps"),
			tgbotapi.NewInlineKeyboardButtonData("Трицепс", "triceps"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ноги", "leg"),
			tgbotapi.NewInlineKeyboardButtonData("Спина", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Плечи", "shoulder"),
		),
	)
	return keyboard
}

func (b *Bot) callbackAllExercisesKeyboard(upd tgbotapi.Update) {
	callback := tgbotapi.NewCallback(upd.CallbackQuery.ID, upd.CallbackQuery.Data)
	if _, err := b.bot.Request(callback); err != nil {
		panic(err)
	}

	var exercise string
	for _, internalSlice := range upd.CallbackQuery.Message.ReplyMarkup.InlineKeyboard {
		for index := range internalSlice {
			if *internalSlice[index].CallbackData == upd.CallbackQuery.Data {
				exercise = internalSlice[index].Text
				if exercise == "Спина" {
					exercise = "Спину" // Подумать как исправить хардкодинг
				}
			}
		}
	}

	chatId := upd.CallbackQuery.Message.Chat.ID
	exerciseMessage := fmt.Sprintf("Присылаю упражнения на <b>%s</b>\nСекундочку...", exercise)
	b.bot.Send(tgbotapi.MessageConfig{BaseChat: tgbotapi.BaseChat{ChatID: chatId},
		ParseMode: "HTML", Text: exerciseMessage})
	time.Sleep(2 * time.Second)

	for _, d := range models.ConnectToDataBase(callback.Text) {
		textInformation := fmt.Sprintf("<b>%s</b>\n\n%s\n\n%s", d.Exercise, d.Description, d.Reference)

		msg := tgbotapi.MessageConfig{BaseChat: tgbotapi.BaseChat{ChatID: chatId}, ParseMode: "HTML",
			DisableWebPagePreview: true, Text: textInformation}
		if _, err := b.bot.Send(msg); err != nil {
			b.bot.Send(tgbotapi.NewMessage(chatId, "Произошла ошибка, уже исправляем"))
			panic(err)
		}
	}
}
