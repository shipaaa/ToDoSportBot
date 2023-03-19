package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	b.gendersUser = make(map[string]string)
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
			if err := b.handleCallback(update); err != nil {
				return
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

func (b *Bot) handleCallback(update tgbotapi.Update) error {
	callbackQuery := b.getCallbackFromKeyboard(update)
	chatID := callbackQuery.Message.Chat.ID
	switch callbackQuery.Data {
	case "man", "woman":
		b.gendersUser[callbackQuery.From.UserName] = callbackQuery.Data
		go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
		go b.sendMessage(chatID, messageAfterGenderSelection)
	case "breast", "biceps", "triceps", "leg", "back", "shoulder":
		exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
		go b.sendMessageForExerciseKeyboard(chatID, exercise)
		go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	case "day1":
		go b.sendKeyboard(chatID, messageMuscleGroupSelectionTrainingCom, b.keyboardTrainingDay1)
		go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	case "day2":
		go b.sendKeyboard(chatID, messageMuscleGroupSelectionTrainingCom, b.keyboardTrainingDay2)
		go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	case "day3":
		go b.sendKeyboard(chatID, messageMuscleGroupSelectionTrainingCom, b.keyboardTrainingDay3)
		go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	case "breastTr1", "bicepsTr1", "tricepsTr1", "legTr1", "backTr1", "shoulderTr1":
		exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
		b.sendMessageForKeyboardTraining(chatID, exercise)
	default:
		log.Panic("there is no such callback")
	}
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	b.sendMessage(message.Chat.ID, messageStartCommand)
	b.handleUsersGender(message)
	return nil
}

func (b *Bot) handleUsersGender(message *tgbotapi.Message) error {
	b.gendersUser[message.From.UserName] = ""
	b.sendKeyboard(message.Chat.ID, messageGenderSelection, b.keyboardSex)
	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, messageHelpCommand)
}

func (b *Bot) handleAllExercisesCommand(message *tgbotapi.Message) error {
	if b.gendersUser[message.From.UserName] == "man" {
		go b.sendKeyboard(message.Chat.ID, messageMuscleGroupSelectionAllEx, b.keyboardAllExercises)
		go b.deleteMessage(message.Chat.ID, message.MessageID)
	} else if b.gendersUser[message.From.UserName] == "woman" {
		b.sendMessage(message.Chat.ID, messageWomanProgram2)
	} else {
		go b.sendMessage(message.Chat.ID, messageGenderDetermination)
		go b.handleUsersGender(message)
	}
	return nil
}

func (b *Bot) handleTrainingProgram(message *tgbotapi.Message) error {
	if b.gendersUser[message.From.UserName] == "man" {
		go b.sendKeyboard(message.Chat.ID, messageSelectDay, b.keyboardTrainingProgram)
		go b.deleteMessage(message.Chat.ID, message.MessageID)
	} else if b.gendersUser[message.From.UserName] == "woman" {
		b.sendMessage(message.Chat.ID, messageWomanProgram1)
	} else {
		go b.sendMessage(message.Chat.ID, messageGenderDetermination)
		go b.handleUsersGender(message)
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, messageUnknownCommand)
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	return b.sendMessage(message.Chat.ID, fmt.Sprintf(messageDefault+"/%s\n/%s\n/%s\n/%s",
		startCommand, helpCommand, allExercisesCommand, trainingCommand))
}
