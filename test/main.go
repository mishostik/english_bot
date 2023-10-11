package main

import (
	"log"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Инициализация токена вашего бота
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	// Установка отладочного режима
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Создание канала для получения обновлений от API
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Бесконечный цикл для обработки сообщений
	for {
		select {
		case update := <-updates:
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
