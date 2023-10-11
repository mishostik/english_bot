package main

import (
	"english_bot/cconstants"
	"english_bot/database"
	"english_bot/models"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

func main() {
	//fmt.Println(tg)

	// DATABASE CONNECT

	URI := ""
	db, err := database.Connect("english_bot", URI)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db.Collection("users"))

	// TELEGRAM CONNECT
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}
	//err = handler.HandleMessages(bot)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic()
	}
	for update := range updates {
		fmt.Printf("update: %v", update)
		switch update.Message.Text {
		case "new user":
			// add new user to user collection
			newUser := models.User{
				UserID:       1,
				Username:     "unicorn",
				RegisteredAt: time.Now(),
				LastActiveAt: time.Now(),
				Level:        "B2",
				Role:         cconstants.RoleAdmin,
			}
			//isRegistered, err := database.UserRepository.RegisterUser(user)

			userRepository := database.UserRepository{}

			// Provide the necessary arguments for the RegisterUser method
			isRegistered, err := userRepository.RegisterUser(&newUser)

			// Handle the error or continue with your code
			if err != nil {
				// Handle the error
				fmt.Println("Error registering the user:", err)
				return
			}
			if isRegistered == false {
				fmt.Println("Registered user: false")
			}
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "feedback")
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic()
			}
		}
	}
}
