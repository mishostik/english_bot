package usecase

import (
	"context"
	constants "english_bot/cconstants"
	incorrectRepository "english_bot/internal/incorrect/repository"
	"english_bot/internal/progress"
	progressRepository "english_bot/internal/progress/repository"
	"english_bot/internal/state"
	"english_bot/internal/task"
	taskRepository "english_bot/internal/task/repository"
	"english_bot/internal/user"
	userRepository "english_bot/internal/user/repository"
	"english_bot/pkg/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

type MessageHandlerUsecase struct {
	userRepo *userRepository.UserRepository
	taskRepo *taskRepository.TaskRepository

	incRepo      *incorrectRepository.IncorrectRepository
	progressRepo *progressRepository.ProgressRepository

	stateUC state.UseCase
}

func NewMessageHandlerUsecase(uRepo *userRepository.UserRepository, tRepo *taskRepository.TaskRepository, incRepo *incorrectRepository.IncorrectRepository, progRepo *progressRepository.ProgressRepository, stateUceCase state.UseCase) MessageHandlerUsecase {
	return MessageHandlerUsecase{
		userRepo:     uRepo,
		taskRepo:     tRepo,
		incRepo:      incRepo,
		progressRepo: progRepo,
		stateUC:      stateUceCase,
	}
}

func (u *MessageHandlerUsecase) GetRandomIncorrectAnswers(incorrectAnswers []string, count int) []string {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	randomGenerator.Shuffle(len(incorrectAnswers), func(i, j int) {
		incorrectAnswers[i], incorrectAnswers[j] = incorrectAnswers[j], incorrectAnswers[i]
	})

	return incorrectAnswers[:count]
}

func (u *MessageHandlerUsecase) GenerateKeyboard(buttons []string) tgbotapi.ReplyKeyboardMarkup {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	randomGenerator.Shuffle(len(buttons), func(i, j int) {
		buttons[i], buttons[j] = buttons[j], buttons[i]
	})

	var keyboardButtons [][]tgbotapi.KeyboardButton
	for _, btn := range buttons {
		row := []tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(btn),
		}
		keyboardButtons = append(keyboardButtons, row)
	}

	return tgbotapi.NewReplyKeyboard(keyboardButtons...)
}

func (u *MessageHandlerUsecase) GetExerciseTranslate(ctx context.Context, userId int, typeId int) (*task.Task, error) {
	var (
		task_ *task.Task
		user_ *user.User

		err error
	)

	user_, err = u.userRepo.UserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	task_, err = u.taskRepo.GetTaskByLevelAndType(ctx, user_.Level, typeId)
	if err != nil {
		return nil, err
	}
	return task_, nil
}

func (u *MessageHandlerUsecase) GetRandomTask(ctx context.Context, userId int) (*task.Task, error) {
	var (
		err   error
		user_ *user.User
		task_ *task.Task
	)
	user_, err = u.userRepo.UserByID(ctx, userId)
	if err != nil {
		log.Println("error getting user")
		return nil, err
	}
	task_, err = u.taskRepo.GetRandomTaskByLevel(ctx, user_.Level)
	if err != nil {
		log.Println("error getting random task")
		return nil, err
	}
	return task_, nil
}

func (u *MessageHandlerUsecase) GetExerciseFillGaps(ctx context.Context, userId int, typeId int) (*task.Task, error) {
	return nil, nil
}

func (u *MessageHandlerUsecase) RegisterUser(ctx context.Context, update tgbotapi.Update) string {
	var messageText string

	userId := update.Message.From.ID
	if userId == 0 {
		return fmt.Sprintf("User id is null")
	}
	temp := user.User{
		UserID:       update.Message.From.ID, // check nil pointers
		Username:     update.Message.From.UserName,
		RegisteredAt: utils.GetMoscowTime(),
		LastActiveAt: utils.GetMoscowTime(),
		Level:        constants.LevelB1,
	}
	userExistence, err := u.userRepo.UserByID(ctx, userId) // cache id
	if err != nil {
		return err.Error()
	}
	if userExistence == nil {
		if err := u.userRepo.RegisterUser(ctx, &temp); err != nil {
			messageText = "Ошибка" // todo: че бля? нахер это юзеру
		} else {
			messageText = constants.TestQuestion
		}
	} else {
		messageText = constants.Continue
	}
	return messageText
}

func (u *MessageHandlerUsecase) Respond(ctx context.Context, update tgbotapi.Update) (string, *task.Task, []string, error) {
	var (
		messageText      string
		buttons          []string
		incorrectAnswers []string = []string{"mock", "mock", "mock"}

		err error

		task_ *task.Task

		allStates []string
	)

	userId := update.Message.From.ID

	if update.Message.Text == constants.DefiniteExercise {
		u.stateUC.RememberUserState(userId, constants.DefiniteExercisesState)

		messageText = "А именно?"
		buttons = []string{constants.MsgTranslateRuToEn, constants.MsgTranslateEnToRu, constants.MsgFillGaps}

		return messageText, nil, buttons, nil // TODO handle this

	} else if update.Message.Text == constants.RandomExercise {
		u.stateUC.RememberUserState(userId, constants.RandomExercisesState)

	} else {
		if update.Message.Text == "fuck" {
			u.stateUC.RememberUserState(userId, constants.MainState)
		}

	}

	allStates = u.stateUC.GetUserState(userId)

	if allStates[len(allStates)-2] == constants.ExercisesState {
		if allStates[len(allStates)-1] == constants.RandomExercisesState {
			// get random task
			task_, err = u.GetRandomTask(ctx, update.Message.From.ID)
			if err != nil {
				return "", nil, nil, fmt.Errorf("error getting the random task: %s", err)
			}
		} else if allStates[len(allStates)-1] == constants.DefiniteExercisesState {

			// проверять последнее сообщение пользователя на то, какой он выбрал тип
			var uState string = ""

			switch update.Message.Text {
			case constants.MsgTranslateEnToRu:
				uState = constants.FromEnToRuExercisesState

			case constants.MsgTranslateRuToEn:
				uState = constants.FromRuToEnExercisesState

			case constants.MsgFillGaps:
				uState = constants.FillGapsExercisesState
			}

			u.stateUC.RememberUserState(userId, uState)

		}

	} else if allStates[len(allStates)-2] == constants.DefiniteExercisesState {

		if allStates[len(allStates)-1] == constants.FromEnToRuExercisesState {
			incorrectAnswers = u.GetRandomIncorrectAnswers(constants.IncorrectAnswersRu, 3)
			task_, err = u.GetExerciseTranslate(ctx, update.Message.From.ID, 2)
			if err != nil {
				return "", nil, nil, nil
			}

		} else if allStates[len(allStates)-1] == constants.FromRuToEnExercisesState {
			incorrectAnswers = u.GetRandomIncorrectAnswers(constants.IncorrectAnswersEn, 3)
			task_, err = u.GetExerciseTranslate(ctx, update.Message.From.ID, 1)
			if err != nil {
				return "", nil, nil, nil
			}

		} else if allStates[len(allStates)-1] == constants.FillGapsExercisesState {
			task_, err = u.GetExerciseFillGaps(ctx, update.Message.From.ID, 3)
			if err != nil {
				return "", nil, nil, err
			}
		}

	} else {
		log.Println("не смогли обработать стейт пользователя")
	}

	if task_ != nil {
		messageText = task_.Question
	} else {
		messageText = constants.TaskNotFoundAnswer
		return messageText, nil, nil, nil
	}

	incorrect, err := u.incRepo.GetAnswers(ctx, task_.ID)
	if err != nil {
		log.Fatalf("error getting incorrect answers: %s", err.Error())
	}

	buttons = []string{task_.Answer, incorrect.A, incorrect.B, incorrect.C}
	for i := 0; i < len(buttons); i++ {
		if buttons[i] == "" && len(incorrectAnswers) > 0 {
			buttons[i] = incorrectAnswers[0]
			incorrectAnswers = incorrectAnswers[1:]
		}
	}

	return messageText, task_, buttons, nil
}

func (u *MessageHandlerUsecase) CheckUserAnswer(ctx context.Context, rightAnswer string, userAnswer string, userId int, taskId uuid.UUID) (string, error) {
	var (
		score uint8
		msg   string
	)
	if rightAnswer == userAnswer {
		score = 2
		msg = constants.MsgRightAnswer
	} else {
		score = 0
		msg = constants.MsgWrongAnswer
	}

	oldProgress, err := u.progressRepo.GetUserProgress(ctx, userId, taskId)
	if err != nil {
		log.Println("error getting old user progress")
		return "", err
	}
	var (
		oldScore      int64  = 0
		oldTaskLevel  string = "" //task.Level // wtf
		receivedTasks int16  = 0
	)
	if oldProgress != nil {
		oldScore = oldProgress.Score
		oldTaskLevel = oldProgress.TaskLevel
		receivedTasks = oldProgress.ReceivedTasks
	}
	res := &progress.UserProgress{
		UserID:        userId,
		Score:         oldScore + int64(score),
		TaskLevel:     oldTaskLevel,
		ReceivedTasks: receivedTasks + 1,
	}

	err = u.progressRepo.InsertUserResult(ctx, res)
	if err != nil {
		return "", err
	}
	return msg, nil
}
