package handlers

import (
	"english_bot/telegram"
	"log"
	"strings"
)

//type MessageHandler interface {
//	HandleMessage(message string, chatID int, username string) error
//}

type MessageHandler struct {
	tg *telegram.Client
}

func (h *MessageHandler) HandleMessage(msg string, chatID int, username string) error {
	msg = strings.TrimSpace(msg)
	log.Printf("get new command '%s' from '%s'", msg, username)
	var response string
	switch msg {
	case "/start":
		response = "hello"
	default:
		response = "unknown"
	}

	err := h.tg.SendMessage(chatID, response)
	if err != nil {
		log.Printf("Error while sending message: %v", err)
		return err
	}
	return nil
}
