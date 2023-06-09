package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func (b *Bot) keyboardSex() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Мужской", "man"),
			tgbotapi.NewInlineKeyboardButtonData("Женский", "woman"),
		),
	)
	return keyboard
}

func (b *Bot) keyboardAllExercises() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Грудь", "breast"),
			tgbotapi.NewInlineKeyboardButtonData("Бицепс", "biceps"),
			tgbotapi.NewInlineKeyboardButtonData("Трицепс", "triceps"),
		),
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ноги", "leg"),
			tgbotapi.NewInlineKeyboardButtonData("Спина", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Плечи", "shoulder"),
		),
	)
	return keyboard
}

func (b *Bot) keyboardTrainingProgram() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("День 1", "day1"),
			tgbotapi.NewInlineKeyboardButtonData("День 2", "day2"),
			tgbotapi.NewInlineKeyboardButtonData("День 3", "day3"),
		),
	)
	return keyboard
}

func (b *Bot) keyboardTrainingDay1() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Грудь", "breastTr1"),
			tgbotapi.NewInlineKeyboardButtonData("Трицепс", "tricepsTr1"),
		),
	)
	return keyboard
}

func (b *Bot) keyboardTrainingDay2() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Спина", "backTr1"),
			tgbotapi.NewInlineKeyboardButtonData("Бицепс", "bicepsTr1"),
		),
	)
	return keyboard
}

func (b *Bot) keyboardTrainingDay3() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		makeInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ноги", "legTr1"),
			tgbotapi.NewInlineKeyboardButtonData("Плечи", "shoulderTr1"),
		),
	)
	return keyboard
}

func makeInlineKeyboardRow(buttons ...tgbotapi.InlineKeyboardButton) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(buttons...)
}

func (b *Bot) getCallbackFromKeyboard(upd tgbotapi.Update) *tgbotapi.CallbackQuery {
	callback := tgbotapi.NewCallback(upd.CallbackQuery.ID, upd.CallbackQuery.Data)
	if _, err := b.api.Request(callback); err != nil {
		log.Error(err)
	}

	return upd.CallbackQuery
}
