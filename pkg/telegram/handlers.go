package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	startCommand        = "start"
	helpCommand         = "help"
	allExercisesCommand = "allexercises"
	trainingCommand     = "training"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					return
				}
				continue
			}

			if err := b.handleMessage(update.Message); err != nil {
				return
			}

		} else if update.CallbackQuery != nil {
			callbackQuery := b.getCallbackFromKeyboard(update)
			switch callbackQuery.Data {
			case "breast", "biceps", "triceps", "leg", "back", "shoulder":
				chatId, exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
				b.sendMessageForExerciseKeyboard(chatId, exercise)
			case "day1", "day2", "day3":
				// something new :)
			}
		}
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStartCommand(message)
	case helpCommand:
		return b.handleHelpCommand(message)
	case allExercisesCommand:
		return b.handleAllExercisesCommand(message)
	case trainingCommand:
		return b.handleTrainingProgram(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "Привет. Ты ввел команду /start")
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "А это команда /help")
}

func (b *Bot) handleAllExercisesCommand(message *tgbotapi.Message) error {
	return b.sendKeyboard(message.Chat.ID, "Выбери какую группу мышц ты хочешь прокачать", keyboardAllExercises)
}

func (b *Bot) handleTrainingProgram(message *tgbotapi.Message) error {
	//return b.sendMessage(message.Chat.ID, "Пока что не работает. Простите :(")
	return b.sendKeyboard(message.Chat.ID, "Выбери день", keyboardTrainingProgram)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "Я не знаю такой команды")
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	return b.sendMessage(message.Chat.ID, fmt.Sprintf("Я еще не настолько умный бот, поэтому общаюсь только"+
		" командами.\nТыкай:\n\n/%s\n/%s\n/%s\n/%s", startCommand, helpCommand, allExercisesCommand, trainingCommand))
}
