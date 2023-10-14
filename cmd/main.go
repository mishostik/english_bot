package main

import (
	"context"
	"english_bot/database"
	handlerUseCase "english_bot/internal/messageHandler/usecase"
	updateUseCase "english_bot/internal/updates/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	ctx := context.Background()
	URI := config.DATABASE_URI
	db, err := database.Connect(ctx, "english_bot", URI)
	if err != nil {
		log.Fatal(err)
	}

	userCollection, err := db.Collection("users")
	if err != nil {
		log.Fatal(err)
	}
	typeCollection, err := db.Collection("task_types")
	if err != nil {
		log.Fatal(err)
	}

	taskCollection, err := db.Collection("tasks")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := database.NewUserRepository(userCollection)
	taskRepo := database.NewTaskRepository(taskCollection)
	typeRepo := database.NewTypeRepository(typeCollection)

	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	updateUC := updateUseCase.Init(bot)
	handler := handlerUseCase.InitHandler(bot, userRepo)
	admin := handlerUseCase.InitAdmin(bot, taskRepo, typeRepo)

	//updates, _ := updateUC.NewUpdates()

	updates, err := bot.GetUpdatesChan(updateUC.UpdateConfig)
	err = admin.HandleTasks(ctx)
	if err != nil {
		return
	}
	err = handler.HandleMessages(ctx, updates)
	if err != nil {
		return
	}

}
