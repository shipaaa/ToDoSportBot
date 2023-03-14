package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
