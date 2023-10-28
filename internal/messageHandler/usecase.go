package messageHandler

import (
	"context"
	"english_bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Usecase interface {
	GetRandomIncorrectAnswers(incorrectAnswers []string, count int) []string
	GenerateKeyboard(buttons []string) tgbotapi.ReplyKeyboardMarkup
	GetExerciseTranslate(ctx context.Context, userID int, typeId int) (*models.Task, error)
	GetRandomTask(ctx context.Context, userId int) (*models.Task, error)
	GetExerciseFillGaps(ctx context.Context, userId int, typeId int) (*models.Task, error)
}
