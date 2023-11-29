package core

import (
	"context"
	constants "english_bot/cconstants"
	"english_bot/internal/dictionary"
	"english_bot/internal/message"
	"english_bot/internal/progress"
	"english_bot/internal/state"
	"english_bot/internal/task"
	"english_bot/internal/user/repository"
	"english_bot/models"
	"english_bot/pkg/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type GeneralHandler struct {
	dictHandler     *dictionary.DictionaryHandler
	exHandler       *task.ExerciseHandler
	progressHandler *progress.ProgressHandler
	messageUC       message.Usecase
	stateUC         state.UseCase
	userRepo        *repository.UserRepository

	task *models.Task
}

func NewGeneralHandler(dHandler *dictionary.DictionaryHandler, eHandler *task.ExerciseHandler, pHandler *progress.ProgressHandler, mUseCase message.Usecase, sUseCase state.UseCase, uRepo *repository.UserRepository) *GeneralHandler {
	return &GeneralHandler{
		dictHandler:     dHandler,
		exHandler:       eHandler,
		progressHandler: pHandler,
		messageUC:       mUseCase,
		stateUC:         sUseCase,
		userRepo:        uRepo,
	}
}

func (h *GeneralHandler) RegisterUser(ctx context.Context, update tgbotapi.Update) string {
	var messageText string

	userId := update.Message.From.ID
	if userId == 0 {
		return fmt.Sprintf("User id is null")
	}
	user := models.User{
		UserID:       update.Message.From.ID, // check nil pointers
		Username:     update.Message.From.UserName,
		RegisteredAt: utils.GetMoscowTime(),
		LastActiveAt: utils.GetMoscowTime(),
		Level:        constants.LevelB1,
	}
	userExistence, err := h.userRepo.UserByID(ctx, userId) // cache id
	if err != nil {
		return err.Error()
	}
	if userExistence == nil {
		if err := h.userRepo.RegisterUser(ctx, &user); err != nil {
			messageText = "Ошибка" // todo: че бля? нахер это юзеру
		} else {
			messageText = constants.TestQuestion
		}
	} else {
		messageText = constants.Continue
	}
	return messageText
}

func (h *GeneralHandler) GetMainMenu(ctx context.Context, update tgbotapi.Update) (string, []string) {
	var (
		messageText string
		buttons     []string
	)
	userId := update.Message.From.ID
	switch update.Message.Text {

	case "/start":
		messageText = h.RegisterUser(ctx, update)
		buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}

	case "тест":
		h.stateUC.RememberUserState(userId, constants.TestState)

		messageText = constants.TestDescription
		buttons = []string{constants.MsgBeginTest, constants.MsgGoBack}

	case "задания":
		h.stateUC.RememberUserState(userId, constants.ExercisesState)

		messageText = constants.ExerciseDescription
		buttons = []string{"любые", "определенные"}

	case "словарь":
		h.stateUC.RememberUserState(userId, constants.DictionaryState)

		messageText = constants.DictionaryDescription
		buttons = []string{constants.MsgAddNewWord, constants.MsgGetContext}

	case "fuck":
		h.stateUC.RememberUserState(userId, constants.MainState)

		messageText = constants.MainStateDescription
		buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}

	default:
		messageText = "mock"
		buttons = []string{"wrong"}
	}
	return messageText, buttons
}

func (h *GeneralHandler) Reply(bot *tgbotapi.BotAPI, ctx context.Context, update tgbotapi.Update) error {
	var (
		messageText = constants.MsgUnknownCommand

		err         error
		buttons     []string
		responseBuf string
	)

	userId := update.Message.From.ID
	userState, err := h.stateUC.GetUserState(userId)
	if err != nil || len(userState) == 0 {
		log.Fatal(err)
	}

	switch userState[len(userState)-1] {

	case constants.ExercisesState:
		messageText, h.task, buttons, err = h.exHandler.Respond(ctx, update)
		log.Println("New task received!!")
		if err != nil {
			return err
		}
		h.stateUC.RememberUserState(userId, constants.ExerciseProcessState)
		//fallthrough

	case constants.ExerciseProcessState:
		log.Println("in exercise ...")
		log.Println("question -", h.task.Question)
		if h.task != nil {

			// здесь message "верно" или "неверно"
			userAnswer := update.Message.Text
			responseBuf, err = h.progressHandler.CheckUserAnswer(ctx, h.task.Answer, userAnswer, userId, h.task.ID)
			if err != nil {
				return err
			}
			h.stateUC.RememberUserState(userId, constants.ExercisesState)

			messageText, h.task, buttons, err = h.exHandler.Respond(ctx, update)
			if err != nil {
				return err
			}
		} else {
			log.Fatal("task not found. task - nil")
			return nil
		}

	case constants.TestState:
		log.Println("test")
		messageText = "Позже здесь появится тест 😏"

	case constants.DictionaryState:
		err = h.dictHandler.Respond(bot, ctx, update)

	case constants.MainState:
		messageText, buttons = h.GetMainMenu(ctx, update)

	default:
		buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}
	}

	if responseBuf != "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseBuf)
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	if len(buttons) > 0 {
		msg.ReplyMarkup = h.messageUC.GenerateKeyboard(buttons)
	}
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
