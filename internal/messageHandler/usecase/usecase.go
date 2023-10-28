package usecase

import (
	"context"
	"english_bot/database"
	"english_bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"time"
)

type MessageHandlerUsecase struct {
	userRepo *database.UserRepository
	taskRepo *database.TaskRepository
}

func NewMessageHandlerUsecase(uRepo *database.UserRepository, tRepo *database.TaskRepository) *MessageHandlerUsecase {
	return &MessageHandlerUsecase{
		userRepo: uRepo,
		taskRepo: tRepo,
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
	for i := 0; i < len(buttons); i += 2 {
		// Создайте два ряда кнопок
		row := []tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(buttons[i]),
		}

		// Проверьте, есть ли следующая кнопка, чтобы избежать выхода за пределы массива
		if i+1 < len(buttons) {
			row = append(row, tgbotapi.NewKeyboardButton(buttons[i+1]))
		}

		keyboardButtons = append(keyboardButtons, row)
	}

	return tgbotapi.NewReplyKeyboard(keyboardButtons...)
}

func (u *MessageHandlerUsecase) GetExerciseTranslate(ctx context.Context, userId int, typeId int) (*models.Task, error) {
	var (
		task *models.Task
		err  error
		user *models.User
	)

	user, err = u.userRepo.UserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	task, err = u.taskRepo.GetTaskByLevelAndType(ctx, user.Level, typeId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *MessageHandlerUsecase) GetRandomTask(ctx context.Context, userId int) (*models.Task, error) {
	var (
		err  error
		user *models.User
		task *models.Task
	)
	user, err = u.userRepo.UserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	task, err = u.taskRepo.GetRandomTaskByLevel(ctx, user.Level)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *MessageHandlerUsecase) GetExerciseFillGaps(ctx context.Context, userId int, typeId int) (*models.Task, error) {
	return nil, nil
}
