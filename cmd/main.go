package main

import (
	"context"
	"english_bot/database"
	handlerUseCase "english_bot/internal/messageHandler/usecase"
	updateUseCase "english_bot/internal/updates/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func main() {
	//fmt.Println(tg)
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	// DATABASE CONNECT
	URI := ""
	db, err := database.Connect(ctx, "english_bot", URI)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := db.Collection("users")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := database.NewUserRepository(collection)

	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Fatal(err)
	}
	updateUC := updateUseCase.Init(bot)
	_ = handlerUseCase.InitHandler(bot, userRepo)

	updates, _ := updateUC.NewUpdates()

	//updates, err := bot.GetUpdatesChan(updateUC.UpdateConfig)

	for {
		select {
		case update := <-*updates:
			if update.Message == nil {
				continue
			}

			log.Printf("[%s] %s", update.Message.From.ID, update.Message.Text)

			// Отправить ответ на полученное сообщение
			reply := "Вы сказали: " + update.Message.Text
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)

		case <-time.After(time.Second * 5):
			// Чтобы избежать проблем с длинными периодами неактивности, бот отправляет пустое сообщение каждые 5 секунд
			// для поддержания активного соединения с серверами Telegram.
			_, _ = bot.Send(tgbotapi.NewMessage(0, ""))
		}
	}

}
