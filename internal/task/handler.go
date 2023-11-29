package task

import (
	"context"
	constants "english_bot/cconstants"
	"english_bot/internal/incorrect/repository"
	"english_bot/internal/message"
	"english_bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type ExerciseHandler struct {
	messageUsecase message.Usecase
	incorrectRepo  *repository.IncorrectRepository
}

func NewExerciseHandler(mUsecase message.Usecase, incRepo *repository.IncorrectRepository) *ExerciseHandler {
	return &ExerciseHandler{
		messageUsecase: mUsecase,
		incorrectRepo:  incRepo,
	}
}

func (h *ExerciseHandler) Respond(ctx context.Context, update tgbotapi.Update) (string, *models.Task, []string, error) {
	var (
		messageText      string
		buttons          []string
		err              error
		incorrectAnswers []string
	)

	if len(update.Message.Text) == 0 {
		return "", nil, nil, nil
	}

	if update.Message.Text == "определенные" {
		messageText = "Выбери тип заданий"
		buttons = []string{constants.MsgTranslateRuToEn, constants.MsgTranslateEnToRu, constants.MsgFillGaps}

	} else {
		incorrectAnswers = []string{"mock", "mock", "mock"}
		switch update.Message.Text {
		// TODO добавить 1 из 3х задачи на уровень +1 или -1 (рандомный уровень)
		case "любые":
			// срабатывает только после получения сообщения "любые" - один раз. добавить второй юзер стейт?
			// TODO изменить логику
			task, err = h.messageUsecase.GetRandomTask(ctx, update.Message.From.ID)
			if err != nil {
				return "", nil, nil, err
			}
		case "перевод на английский":
			incorrectAnswers = h.messageUsecase.GetRandomIncorrectAnswers(constants.IncorrectAnswersEn, 3)
			task, err = h.messageUsecase.GetExerciseTranslate(ctx, update.Message.From.ID, 1)
			if err != nil {
				return "", nil, nil, err
			}

		case "перевод на русский":
			incorrectAnswers = h.messageUsecase.GetRandomIncorrectAnswers(constants.IncorrectAnswersRu, 3)
			task, err = h.messageUsecase.GetExerciseTranslate(ctx, update.Message.From.ID, 2)
			if err != nil {
				return "", nil, nil, err
			}

		case "заполнить пропуски":
			task, err = h.messageUsecase.GetExerciseFillGaps(ctx, update.Message.From.ID, 3)
			if err != nil {
				return "", nil, nil, err
			}

		default:
			messageText = "не получилось найти задачу 😢"
		}
		if task != nil {
			messageText = task.Question
		} else {
			messageText = "задачи нет 😱"
		}

		// get incorrect answers from db
		incorrect, err := h.incorrectRepo.GetAnswers(ctx, task.ID)
		if err != nil {
			log.Println("error getting incorrect answers")
		}

		buttons = []string{task.Answer, incorrect.A, incorrect.B, incorrect.C}
		for i := 0; i < len(buttons); i++ {
			if buttons[i] == "" && len(incorrectAnswers) > 0 {
				buttons[i] = incorrectAnswers[0]
				incorrectAnswers = incorrectAnswers[1:]
			}
		}
	}

	return messageText, task, buttons, nil
}
