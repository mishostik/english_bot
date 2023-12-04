package message

import (
	"context"
	constants "english_bot/cconstants"
	"english_bot/internal/dictionary"
	"english_bot/internal/message/usecase"
	"english_bot/internal/state"
	"english_bot/internal/task"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type MessageHandler struct {
	bot     *tgbotapi.BotAPI
	stateUC state.UseCase
	task    *task.Task

	dictH *dictionary.DictionaryHandler
	msgUC usecase.MessageHandlerUsecase
}

func InitHandler(bot *tgbotapi.BotAPI, stateUseCase state.UseCase, dHandler *dictionary.DictionaryHandler, messageUC usecase.MessageHandlerUsecase) *MessageHandler {
	return &MessageHandler{
		bot:     bot,
		stateUC: stateUseCase,
		dictH:   dHandler,
		msgUC:   messageUC,
	}
}

func (h *MessageHandler) GetMainMenu(ctx context.Context, update tgbotapi.Update) (string, []string) {
	var (
		messageText string
		buttons     []string
	)
	userId := update.Message.From.ID
	switch update.Message.Text {

	case "/start":
		messageText = h.msgUC.RegisterUser(ctx, update)
		buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}

	case constants.MsgTest:
		h.stateUC.RememberUserState(userId, constants.TestState)

		messageText = constants.TestDescription
		buttons = []string{constants.MsgBeginTest, constants.MsgGoBack}

	case constants.MsgTasks:
		h.stateUC.RememberUserState(userId, constants.ExercisesState)

		messageText = constants.ExerciseDescription
		buttons = []string{constants.RandomExercise, constants.DefiniteExercise}

	case constants.MsgDictionary:
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

func (h *MessageHandler) Reply(bot *tgbotapi.BotAPI, ctx context.Context, update tgbotapi.Update) error {
	var (
		messageText = constants.MsgUnknownCommand

		err         error
		buttons     []string
		responseBuf string
	)

	userId := update.Message.From.ID
	userState := h.stateUC.GetUserState(userId)
	if len(userState) == 0 {
		h.stateUC.RememberUserState(userId, constants.MainState)
		messageText, buttons = h.GetMainMenu(ctx, update)
	}

	if len(userState) > 1 {

		// TODO просчитывать на основании последний нескольких стейтов
		switch userState[len(userState)-1] {

		case constants.ExercisesState:
			messageText, h.task, buttons, err = h.msgUC.Respond(ctx, update)
			if err != nil {
				return err
			}
			if h.task != nil {
				h.stateUC.RememberUserState(userId, constants.ExerciseProcessState)
			}
			//fallthrough

		case constants.ExerciseProcessState:

			if h.task != nil {
				userAnswer := update.Message.Text
				responseBuf, err = h.msgUC.CheckUserAnswer(ctx, h.task.Answer, userAnswer, userId, h.task.ID)
				if err != nil {
					return err
				}

				h.stateUC.RememberUserState(userId, constants.ExercisesState)

				messageText, h.task, buttons, err = h.msgUC.Respond(ctx, update)
				if err != nil {
					return err
				}
			} else {
				log.Fatal("task not found. task - nil")
				return nil
			}

		case constants.TestState:
			messageText = "Позже здесь появится тест 😏"

		case constants.DictionaryState:
			err = h.dictH.Respond(bot, ctx, update)

		case constants.MainState:
			messageText, buttons = h.GetMainMenu(ctx, update)

		default:
			// ???
			buttons = []string{constants.MsgDictionary, constants.MsgTasks, constants.MsgTest}
		}
	} else {
		messageText, buttons = h.GetMainMenu(ctx, update)
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
		msg.ReplyMarkup = h.msgUC.GenerateKeyboard(buttons)
	}
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) HandleMessages(ctx context.Context, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		fmt.Println(update)
		if err := h.Reply(h.bot, ctx, update); err != nil {
			log.Println(fmt.Sprintf("error in reply {%s}", err.Error()))
		}
	}
	return nil
}
