package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
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
			chatID := callbackQuery.Message.Chat.ID
			switch callbackQuery.Data {
			case "breast", "biceps", "triceps", "leg", "back", "shoulder":
				exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
				time.Sleep(time.Second)
				go b.sendMessageForExerciseKeyboard(chatID, exercise)
				go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
			case "day1":
				go b.sendKeyboard(chatID, "Выбери группу мышц", b.keyboardTrainingDay1)
				go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
			case "day2":
				go b.sendKeyboard(chatID, "Выбери группу мышц", b.keyboardTrainingDay2)
				go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
			case "day3":
				go b.sendKeyboard(chatID, "Выбери группу мышц", b.keyboardTrainingDay3)
				go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
			case "breastTr1", "bicepsTr1", "tricepsTr1", "legTr1", "backTr1", "shoulderTr1":
				exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
				time.Sleep(time.Second)
				b.sendMessageForKeyboardTraining(chatID, exercise)
				//go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
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
	b.sendMessage(message.Chat.ID, "Привееет!!!")
	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "А это команда /help")
}

func (b *Bot) handleAllExercisesCommand(message *tgbotapi.Message) error {
	go b.sendKeyboard(message.Chat.ID, "Выбери какую группу мышц ты хочешь прокачать", b.keyboardAllExercises)
	go b.deleteMessage(message.Chat.ID, message.MessageID)
	return nil
}

func (b *Bot) handleTrainingProgram(message *tgbotapi.Message) error {
	go b.sendKeyboard(message.Chat.ID, "Выбери день", b.keyboardTrainingProgram)
	go b.deleteMessage(message.Chat.ID, message.MessageID)
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "Я не знаю такой команды")
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	return b.sendMessage(message.Chat.ID, fmt.Sprintf("Я еще не настолько умный бот, поэтому общаюсь только"+
		" командами.\nТыкай:\n\n/%s\n/%s\n/%s\n/%s", startCommand, helpCommand, allExercisesCommand, trainingCommand))
}
