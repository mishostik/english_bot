package usecase

import (
	"context"
	"english_bot/handlers"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type MessageHandler struct {
	bot      *tgbotapi.BotAPI
	gHandler *handlers.GeneralHandler
}

func InitHandler(bot *tgbotapi.BotAPI, gHandler *handlers.GeneralHandler) *MessageHandler {
	return &MessageHandler{
		bot:      bot,
		gHandler: gHandler,
	}
}

func (h *MessageHandler) HandleMessages(ctx context.Context, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		fmt.Println(update)
		if err := h.gHandler.Reply(h.bot, ctx, update); err != nil {
			log.Println(fmt.Sprintf("error in reply {%s}", err.Error()))
		}
	}
	return nil
}
