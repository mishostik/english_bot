package usecase

import (
	"context"
	"english_bot/constants"
	"english_bot/database"
	"english_bot/models"
	"english_bot/pkg/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type MessageHandler struct {
	bot      *tgbotapi.BotAPI
	userRepo *database.UserRepository
}

func InitHandler(bot *tgbotapi.BotAPI, userRepo *database.UserRepository) *MessageHandler {
	return &MessageHandler{
		bot:      bot,
		userRepo: userRepo,
	}
}

func (h *MessageHandler) GenerateKeyboard(buttons []string) tgbotapi.ReplyKeyboardMarkup {
	var keyboardButtons []tgbotapi.KeyboardButton
	for _, button := range buttons {
		keyboardButtons = append(keyboardButtons, tgbotapi.NewKeyboardButton(button))
	}
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(keyboardButtons...),
	)
}

func (h *MessageHandler) RegisterUser(update tgbotapi.Update, ctx context.Context) string {
	var messageText = "..."
	userID := update.Message.From.ID
	user := models.User{
		UserID:       update.Message.From.ID,
		Username:     update.Message.From.UserName,
		RegisteredAt: utils.GetMoscowTime(),
		LastActiveAt: utils.GetMoscowTime(),
		Level:        constants.LevelA0,
		Role:         constants.RoleUser,
	}
	userExistence, err := h.userRepo.UserByID(userID)
	if err != nil {
		messageText = "error while finding user"
	}
	if userExistence == nil {
		if err := h.userRepo.RegisterUser(ctx, &user); err != nil {
			messageText = "error while registration"
		} else {
			messageText = "хотели бы пройти тест для определения вашего уровня владения английским языком?"
			// предлагать пройти тест
		}
	} else {
		messageText = "user already exist"
	}
	return messageText
}

func (h *MessageHandler) Reply(ctx context.Context, update tgbotapi.Update) error {
	var (
		user        models.User
		messageText = constants.MsgUnknownCommand
	)
	fmt.Println(user)
	buttons := []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}

	switch update.Message.Text {
	case "/start":
		messageText = h.RegisterUser(update, ctx)
		buttons = []string{constants.MsgYes, constants.MsgNo}
	case "тест":
		messageText = "...Информация о тесте..."
		buttons = []string{constants.MsgBeginTest, constants.MsgGoBack}
	case "задания":
		messageText = "Выбери тип заданий"
		buttons = []string{constants.MsgTranslateRuToEn, constants.MsgTranslateEnToRu, constants.MsgFillGaps}
	case "словарь":
		messageText = "Сюда можно добавить новые слова"
		buttons = []string{constants.MsgAddNewWord, constants.MsgGetContext}
	default:
		messageText = "default"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	msg.ReplyMarkup = h.GenerateKeyboard(buttons)
	_, err := h.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) HandleMessages(ctx context.Context, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		fmt.Println(update)
		if err := h.Reply(ctx, update); err != nil {
			log.Println(fmt.Sprintf("error in reply {%s}", err.Error()))
		}
	}
	return nil
}
