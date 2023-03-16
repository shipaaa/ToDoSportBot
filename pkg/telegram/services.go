package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getExerciseName(k *tgbotapi.InlineKeyboardMarkup, cbq *tgbotapi.CallbackQuery) string {
	var exercise string
	for _, v := range k.InlineKeyboard {
		for i := range v {
			if *v[i].CallbackData == cbq.Data {
				exercise = v[i].Text
				if exercise == "Спина" {
					exercise = "Спину" // Подумать как исправить хардкодинг
				}
			}
		}
	}
	return exercise
}
