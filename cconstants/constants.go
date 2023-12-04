package constants

const (
	RoleAdmin = "Admin"
	RoleUser  = "User"
)

const (
	LevelA0 = "A0"
	LevelA1 = "A1"
	LevelA2 = "A2"
	LevelB1 = "B1"
	LevelB2 = "B2"
	LevelC1 = "C1"
)

// TODO раписать логику стейта пользователя
const (
	MainState            = "Main"
	TestState            = "Test"
	ExercisesState       = "Exercises"
	ExerciseProcessState = "ExerciseProcess"

	DefiniteExercisesState   = "DefiniteExercises"
	RandomExercisesState     = "RandomExercises"
	FromRuToEnExercisesState = "FromRuToEnExercises"
	FromEnToRuExercisesState = "FromEnToRuExercises"
	FillGapsExercisesState   = "FillGapsExercises"
	DictionaryState          = "Dictionary"
	TranslatorState          = "Translator"
)

// TODO убрать это нахуй отсюда

var IncorrectAnswersEn = []string{"Exam", "Condition", "Blue", "Exercise", "Forest", "Space", "Rain", "Father", "Beast"}
var IncorrectAnswersRu = []string{"Решение", "Космос", "Дождь", "Отец", "Зверь", "Экзамен", "Синий", "Лес", "Пример"}

const (
	TestDescription       = "Тест на определения уровня владения языком"
	ExerciseDescription   = "Выбери тип заданий"
	DictionaryDescription = "Сюда можно добавить новые слова"
	MainStateDescription  = "Чем займемся?"
	TestQuestion          = "Хотели бы пройти тест для определения вашего уровня владения английским языком?"
	Continue              = "Продолжим"

	//---- Exercises ------------------------------------------------------------------------------------------------------------

	DefiniteExercise = "Определенные"
	RandomExercise   = "Любые"
)
