package message

import (
	"context"
	"english_bot/internal/task"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Usecase interface {
	GetRandomIncorrectAnswers(incorrectAnswers []string, count int) []string
	GenerateKeyboard(buttons []string) tgbotapi.ReplyKeyboardMarkup
	GetExerciseTranslate(ctx context.Context, userID int, typeId int) (*task.Task, error)
	GetRandomTask(ctx context.Context, userId int) (*task.Task, error)
	GetExerciseFillGaps(ctx context.Context, userId int, typeId int) (*task.Task, error)
	RegisterUser(ctx context.Context, update tgbotapi.Update) string
	Respond(ctx context.Context, update tgbotapi.Update) (string, *task.Task, []string, error)
}
