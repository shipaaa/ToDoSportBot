package telegram

const (
	startCommand        = "start"
	helpCommand         = "help"
	allExercisesCommand = "allexercises"
	trainingCommand     = "training"
)

const (
	msgDefault        = "Я еще не настолько умный бот, поэтому общаюсь только командами.\nТыкай:\n\n"
	msgStartCommand   = "Привееет!!! Для начала определимся кто ты ⬇️"
	msgUnknownCommand = "Я не знаю такой команды 🙁"
	msgHelpCommand    = "Если хочешь посмотреть чужие упражнения, начни заново и выбери другой пол — /start\n" +
		"Хочешь посмотреть все упражнения, которые есть в этом боте — /allexercises\n" +
		"Если же хочешь посмотреть что-то конкретное, например, программу тренировок — /training\n" +
		"Хочешь предложить какую то идею или добавить упражнение — пиши: <a href=\"https://t.me/almost_shipa\">Мой телеграм</a>\n" +
		"Если найдешь какой-либо баг или готов усовершенствовать бота в чём-то другом — кидай pull request сюда: <a href=\"https://github.com/shipaaa/ToDoSportBot\">Ссылка на репозиторий</a>"
	msgGenderDetermination             = "Для начала давай определимся с полом"
	msgGenderSelection                 = "Выбери пол:"
	msgAfterGenderSelection            = "Хорошо! Теперь можешь посмотреть все упражнения, либо выбрать программу тренировок"
	msgMuscleGroupSelectionTrainingCom = "Выбери группу мышц"
	msgMuscleGroupSelectionAllEx       = "Выбери какую группу мышц ты хочешь прокачать"
	msgAboutSendEx                     = "Присылаю упражнения на <b>%s</b>\nСекундочку..."
	msgSelectDay                       = "Выбери день"
	msgWomanProgram1                   = "Уже очень скоро здесь будет программа тренировок для девушек :*"
	msgWomanProgram2                   = "Скоро добавим упражнения для прекрасных дам :)"
)
