package usecase

import (
	"context"
	"english_bot/internal/task/repository"
	progress2 "english_bot/internal/user/repository"
	"english_bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"time"
)

type MessageHandlerUsecase struct {
	userRepo *progress2.UserRepository
	taskRepo *repository.TaskRepository
}

func NewMessageHandlerUsecase(uRepo *progress2.UserRepository, tRepo *repository.TaskRepository) *MessageHandlerUsecase {
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
	for _, btn := range buttons {
		row := []tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(btn),
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
		log.Println("error getting user")
		return nil, err
	}
	task, err = u.taskRepo.GetRandomTaskByLevel(ctx, user.Level)
	if err != nil {
		log.Println("error getting random task")
		return nil, err
	}
	return task, nil
}

func (u *MessageHandlerUsecase) GetExerciseFillGaps(ctx context.Context, userId int, typeId int) (*models.Task, error) {
	return nil, nil
}
