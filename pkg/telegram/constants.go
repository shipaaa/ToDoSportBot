package telegram

const (
	startCommand        = "start"
	helpCommand         = "help"
	allExercisesCommand = "allexercises"
	trainingCommand     = "training"
)

const (
	msgDefault        = "Я еще не настолько умный бот, поэтому общаюсь только командами.\nТыкай:\n\n/start\n/help\n/allexercises\n/training"
	msgStartCommand   = "Рад тебя видеть!!"
	msgUnknownCommand = "Я не знаю такой команды 🙁"
	msgHelpCommand    = "Посмотреть упражнения другого пола  —\n/start\n" +
		"Все упражнения в этом боте — /allexercises\n" +
		"Программы тренировок — /training\n" +
		"Пришли мне свои персональные упражнения и я добавлю их для тебя — <a href=\"https://t.me/almost_shipa\">Мой телеграм</a>\n" +
		"Нашёл какой-либо баг или готов усовершенствовать бота в чём-то другом? Кидай pull request сюда — <a href=\"https://github.com/shipaaa/ToDoSportBot\">Ссылка на репозиторий</a>"
	msgGenderDetermination  = "Необходимо выбрать пол"
	msgGenderSelection      = "Выбери пол ⬇️"
	msgAfterGenderSelection = "Теперь выбери подходящий блок в разделе меню"
	msgMuscleGroupSelection = "Выбери нужную группу мышц ⬇️"
	msgAboutSendEx          = "Упражнения на <b>%s</b> 📍"
	msgSelectDay            = "Выбери день ⬇️"
	msgWomanProgram1        = "Уже очень скоро здесь будет программа тренировок для девушек :*"
	msgWomanProgram2        = "Скоро добавим упражнения для прекрасных дам :)"
)
