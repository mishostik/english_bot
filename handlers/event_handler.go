package handlers

import (
	"english_bot/models"
	"english_bot/telegram/types"
)

func event(upd types.Update) (string, models.User, int) {
	msgText := fetchText(upd)
	username := upd.Message.From
	chatId := upd.ChatID
	return msgText, username, chatId
}

func fetchText(upd types.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
