package handlers

import (
	"context"
	constants "english_bot/cconstants"
	"english_bot/database"
	"english_bot/internal/messageHandler"
	"english_bot/models"
	"english_bot/pkg/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var userState = "main"

type GeneralHandler struct {
	dictHandler     *DictionaryHandler
	exHandler       *ExerciseHandler
	progressHandler *ProgressHandler
	messageUsecase  messageHandler.Usecase
	userRepo        *database.UserRepository
}

func NewGeneralHandler(dHandler *DictionaryHandler, eHandler *ExerciseHandler, pHandler *ProgressHandler, mUsecase messageHandler.Usecase, uRepo *database.UserRepository) *GeneralHandler {
	return &GeneralHandler{
		dictHandler:     dHandler,
		exHandler:       eHandler,
		progressHandler: pHandler,
		messageUsecase:  mUsecase,
		userRepo:        uRepo,
	}
}

func (h *GeneralHandler) RegisterUser(ctx context.Context, update tgbotapi.Update) string {
	var messageText string

	userId := update.Message.From.ID
	user := models.User{
		UserID:       update.Message.From.ID, // check nil pointers
		Username:     update.Message.From.UserName,
		RegisteredAt: utils.GetMoscowTime(),
		LastActiveAt: utils.GetMoscowTime(),
		Level:        constants.LevelA0,
		// Role:         constants.RoleUser,
	}
	userExistence, err := h.userRepo.UserByID(ctx, userId) // cache id
	if err != nil {
		return err.Error()
	}
	if userExistence == nil {
		if err := h.userRepo.RegisterUser(ctx, &user); err != nil {
			messageText = "error while registration" // TODO нах пользователю это возвращать ?
		} else {
			messageText = "хотели бы пройти тест для определения вашего уровня владения английским языком?"
		}
	} else {
		messageText = "user already exist"
	}
	return messageText
}

func (h *GeneralHandler) GetMainMenu(ctx context.Context, update tgbotapi.Update) (string, []string) {
	var (
		messageText string
		buttons     []string
	)
	switch update.Message.Text {
	case "/start":
		messageText = h.RegisterUser(ctx, update)
		buttons = []string{constants.MsgYes, constants.MsgNo}
	case "тест":
		userState = "test"
		messageText = "Тест на определения уровня владения языком"
		buttons = []string{constants.MsgBeginTest, constants.MsgGoBack}
	case "задания":
		userState = "exercise"
		messageText = "Выбери тип заданий"
		buttons = []string{"любые", "определенные"}
	case "словарь":
		userState = "dictionary"
		messageText = "Сюда можно добавить новые слова"
		buttons = []string{constants.MsgAddNewWord, constants.MsgGetContext}

	case "на главную":
		userState = "main"
		messageText = "Чем займемся?"
		buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}

	case constants.MsgGoBack:
		userState = "main"
		messageText = "Чем займемся?"
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
	)

	switch userState {

	case "exercise":
		messageText, buttons, err = h.exHandler.Respond(ctx, update)
		if err != nil {
			return err
		}

	case "test":
		log.Println("")

	case "dictionary":
		err = h.dictHandler.Respond(bot, ctx, update)

	case "main":
		messageText, buttons = h.GetMainMenu(ctx, update)

	default:
		buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	msg.ReplyMarkup = h.messageUsecase.GenerateKeyboard(buttons)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
