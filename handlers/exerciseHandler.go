package handlers

import (
	"context"
	constants "english_bot/cconstants"
	"english_bot/internal/messageHandler"
	"english_bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var incorrectAnswersEn = []string{"Exam", "Condition", "Blue", "Exercise", "Forest", "Space", "Rain", "Father", "Beast"}

type ExerciseHandler struct {
	messageUsecase messageHandler.Usecase
}

func NewExerciseHandler(mUsecase messageHandler.Usecase) *ExerciseHandler {
	return &ExerciseHandler{
		messageUsecase: mUsecase,
	}
}

func (h *ExerciseHandler) Respond(ctx context.Context, update tgbotapi.Update) (string, []string, error) {
	var (
		messageText string
		buttons     []string
		task        *models.Task
		err         error
	)

	if len(update.Message.Text) == 0 {
		return "", nil, nil
	}

	if update.Message.Text == "определенные" {
		messageText = "Выбери тип заданий"
		buttons = []string{constants.MsgTranslateRuToEn, constants.MsgTranslateEnToRu, constants.MsgFillGaps}

	} else {

		switch update.Message.Text {

		case "любые":
			task, err = h.messageUsecase.GetRandomTask(ctx, update.Message.From.ID)
			if err != nil {
				return "", nil, err
			}

		case "перевод на английский":
			task, err = h.messageUsecase.GetExerciseTranslate(ctx, update.Message.From.ID, 1)
			if err != nil {
				return "", nil, err
			}

		case "перевод на русский":
			task, err = h.messageUsecase.GetExerciseTranslate(ctx, update.Message.From.ID, 2)
			if err != nil {
				return "", nil, err
			}

		case "заполнить пропуски":
			task, err = h.messageUsecase.GetExerciseFillGaps(ctx, update.Message.From.ID, 3)
			if err != nil {
				return "", nil, err
			}
		}

		messageText = task.Question
		incorrectAnswers := h.messageUsecase.GetRandomIncorrectAnswers(incorrectAnswersEn, 3)

		buttons = []string{task.Answer, "", "", ""}
		for i := 0; i < len(buttons); i++ {
			if buttons[i] == "" && len(incorrectAnswers) > 0 {
				buttons[i] = incorrectAnswers[0]
				incorrectAnswers = incorrectAnswers[1:]
			}
		}
	}

	return messageText, buttons, nil
}
