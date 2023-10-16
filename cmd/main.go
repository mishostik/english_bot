package main

import (
	"context"
	"english_bot/config"
	"english_bot/database"
	handlerUseCase "english_bot/internal/messageHandler/usecase"
	updateUseCase "english_bot/internal/updates/usecase"
	"english_bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {

	ctx := context.Background()

	//TODO: add AE_KEY and encrypt config
	//var shield = secure.NewShield(os.Getenv("AE_KEY"))

	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	//TODO: add AE_KEY and encrypt config
	//config.DecryptConfig(cfg, shield)

	logger := logger.NewLogger(cfg).Sugar()
	defer logger.Sync()

	//TODO: вынести в отдельный файл инициализацию бд

	db, err := database.Connect(ctx, "english_bot", cfg)
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

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
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
