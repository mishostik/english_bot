package usecase

import (
	"context"
	"english_bot/cconstants"
	"english_bot/database"
	"english_bot/models"
	"english_bot/pkg/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
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

func (h *MessageHandler) Reply(ctx context.Context, update tgbotapi.Update) error {
	switch {
	case update.Message.Text == "/start":
		var (
			user        models.User
			messageText = "пошел нахуй"
		)

		user = models.User{
			UserID:       update.Message.From.ID,
			Username:     update.Message.From.UserName,
			RegisteredAt: utils.GetMoscowTime(),
			LastActiveAt: utils.GetMoscowTime(),
			Level:        cconstants.LevelA0,
			Role:         cconstants.RoleUser,
		}

		if err := h.userRepo.RegisterUser(ctx, &user); err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
			h.bot.Send(msg)
			return err
		}
		messageText = "success"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		h.bot.Send(msg)
	case update.Message == nil:
		return fmt.Errorf("message is nil")
	}

	return nil
}

func (h *MessageHandler) HandleMessages(ctx context.Context, updates tgbotapi.UpdatesChannel) error {
	select {
	case update := <-updates:
		if update.Message == nil {
			return fmt.Errorf("message is nil")
		}
		if err := h.Reply(ctx, update); err != nil {
			log.Println(fmt.Sprintf("error in reply {%s}", err.Error()))
		}
	case <-time.After(time.Second * 1):
	}
	//for _, update := range updates {
	//	if err := h.Reply(ctx, update); err != nil {
	//		log.Println(fmt.Sprintf("error in reply {%s}", err.Error()))
	//	}
	//}
	return nil
}
