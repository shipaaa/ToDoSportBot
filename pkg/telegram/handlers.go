package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	b.gendersUser = make(map[string]string)

	for update := range updates {
		if update.Message != nil {
			b.handleMessageOrCommand(update.Message)
		} else if update.CallbackQuery != nil {
			if err := b.handleCallback(update); err != nil {
				log.Error(err)
				return
			}
		}
	}
}

func (b *Bot) handleMessageOrCommand(msg *tgbotapi.Message) {
	var err error
	if msg.IsCommand() {
		err = b.handleCommand(msg)
	} else {
		err = b.handleMessage(msg)
	}

	if err != nil {
		log.Error(err)
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
		err := b.handleGenderSelection(callbackQuery)
		if err != nil {
			return err
		}
	case "breast", "biceps", "triceps", "leg", "back", "shoulder":
		err := b.handleMuscleGroupSelection(callbackQuery)
		if err != nil {
			return err
		}
	case "day1", "day2", "day3":
		err := b.handleTrainingDaySelection(chatID, callbackQuery)
		if err != nil {
			return err
		}
	case "breastTr1", "bicepsTr1", "tricepsTr1", "legTr1", "backTr1", "shoulderTr1":
		exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
		b.sendMessageForKeyboardTraining(chatID, exercise)
	default:
		log.Warning("There is no such callback")
	}
	return nil
}

func (b *Bot) handleGenderSelection(callbackQuery *tgbotapi.CallbackQuery) error {
	b.gendersUser[callbackQuery.From.UserName] = callbackQuery.Data
	chatID := callbackQuery.Message.Chat.ID
	go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	go b.sendMessage(chatID, msgAfterGenderSelection)
	return nil
}

func (b *Bot) handleMuscleGroupSelection(callbackQuery *tgbotapi.CallbackQuery) error {
	exercise := b.sendWaitingMessage(callbackQuery.Message.ReplyMarkup, callbackQuery)
	chatID := callbackQuery.Message.Chat.ID
	go b.sendMessageForExerciseKeyboard(chatID, exercise)
	go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	return nil
}

func (b *Bot) handleTrainingDaySelection(chatID int64, callbackQuery *tgbotapi.CallbackQuery) error {
	switch callbackQuery.Data {
	case "day1":
		go b.sendKeyboard(chatID, msgMuscleGroupSelection, b.keyboardTrainingDay1)
	case "day2":
		go b.sendKeyboard(chatID, msgMuscleGroupSelection, b.keyboardTrainingDay2)
	case "day3":
		go b.sendKeyboard(chatID, msgMuscleGroupSelection, b.keyboardTrainingDay3)
	}
	go b.deleteMessage(chatID, callbackQuery.Message.MessageID)
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	if b.gendersUser[message.From.UserName] == "" {
		b.sendMessage(message.Chat.ID, msgStartCommand)
	}
	b.handleUsersGender(message)
	return nil
}

func (b *Bot) handleUsersGender(message *tgbotapi.Message) error {
	b.gendersUser[message.From.UserName] = ""
	b.sendKeyboard(message.Chat.ID, msgGenderSelection, b.keyboardSex)
	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, msgHelpCommand)
}

func (b *Bot) handleAllExercisesCommand(message *tgbotapi.Message) error {
	if b.gendersUser[message.From.UserName] == "man" {
		go b.sendKeyboard(message.Chat.ID, msgMuscleGroupSelection, b.keyboardAllExercises)
		go b.deleteMessage(message.Chat.ID, message.MessageID)
	} else if b.gendersUser[message.From.UserName] == "woman" {
		b.sendMessage(message.Chat.ID, msgWomanProgram1)
	} else {
		b.sendMessage(message.Chat.ID, msgGenderDetermination)
		b.handleUsersGender(message)
	}
	return nil
}

func (b *Bot) handleTrainingProgram(message *tgbotapi.Message) error {
	if b.gendersUser[message.From.UserName] == "man" {
		go b.sendKeyboard(message.Chat.ID, msgSelectDay, b.keyboardTrainingProgram)
		go b.deleteMessage(message.Chat.ID, message.MessageID)
	} else if b.gendersUser[message.From.UserName] == "woman" {
		b.sendMessage(message.Chat.ID, msgWomanProgram2)
	} else {
		b.sendMessage(message.Chat.ID, msgGenderDetermination)
		b.handleUsersGender(message)
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, msgUnknownCommand)
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Infof("[%s] %s", message.From.UserName, message.Text)
	return b.sendMessage(message.Chat.ID, msgDefault)
}
