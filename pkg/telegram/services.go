package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"strings"
)

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

func getMessageText(table []models.Table) string {
	var sl []string
	for _, t := range table {
		dbRow := fmt.Sprintf("<b>%d. %s</b>\n\n%s\n\n<a href=\"%s\">Видосик обучалка</a>",
			t.Id, t.Exercise, t.Description, t.Reference)
		sl = append(sl, dbRow)
	}
	messageText := strings.Join(sl, "\n\n")
	return messageText
}
