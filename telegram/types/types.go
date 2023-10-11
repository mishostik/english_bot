package types

import "english_bot/models"

type Message struct {
	MessageID int         `json:"message_id"`
	From      models.User `json:"from"`
	Text      string      `json:"text"`
}

type Update struct {
	UpdateID int `json:"update_id"`
	ChatID   int `json:"chat_id"`
	Message  *Message
}
