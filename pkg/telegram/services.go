package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shipaaa/telegram-sport-bot/pkg/models"
	"strings"
)

func getExerciseName(k *tgbotapi.InlineKeyboardMarkup, cbq *tgbotapi.CallbackQuery) string {
	button := getSelectedButton(k, cbq)
	if button != nil {
		exercise := button.Text
		exercise = fixName(exercise)
		return exercise
	}
	return ""
}

func getSelectedButton(k *tgbotapi.InlineKeyboardMarkup,
	cbq *tgbotapi.CallbackQuery) *tgbotapi.InlineKeyboardButton {

	for _, row := range k.InlineKeyboard {
		for _, button := range row {
			if button.CallbackData != nil && *button.CallbackData == cbq.Data {
				return &button
			}
		}
	}
	return nil
}

func fixName(exercise string) string {
	switch exercise {
	case "Спина":
		exercise = "Спину"
	}
	return exercise
}

func getMessageTextForExerciseKeyboard(table []models.Table) string {
	var rows []string
	for _, t := range table {
		row := getTableRowText(t)
		rows = append(rows, row)
	}
	return strings.Join(rows, "\n\n")
}

func getTableRowText(t models.Table) string {
	row := fmt.Sprintf("<b>%d. %s</b>\n\n%s\n\n<a href=\"%s\">Видосик обучалка</a>",
		t.Id, t.Exercise, t.Description, t.Reference)
	return row
}
