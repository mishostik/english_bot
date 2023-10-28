package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DictionaryHandler struct {
}

func NewDictionaryHandler() *DictionaryHandler {
	return &DictionaryHandler{}
}

func (h *DictionaryHandler) Respond(bot *tgbotapi.BotAPI, ctx context.Context, update tgbotapi.Update) error {
	return nil
}
