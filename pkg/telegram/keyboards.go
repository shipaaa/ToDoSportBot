package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func keyboardAllExercises() tgbotapi.InlineKeyboardMarkup {
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

func keyboardTrainingProgram() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("День 1", "day1"),
			tgbotapi.NewInlineKeyboardButtonData("День 2", "day2"),
			tgbotapi.NewInlineKeyboardButtonData("День 3", "day3"),
		),
	)
	return keyboard
}

func (b *Bot) getCallbackFromKeyboard(upd tgbotapi.Update) *tgbotapi.CallbackQuery {
	callback := tgbotapi.NewCallback(upd.CallbackQuery.ID, upd.CallbackQuery.Data)
	if _, err := b.api.Request(callback); err != nil {
		panic(err)
	}

	return upd.CallbackQuery
}
