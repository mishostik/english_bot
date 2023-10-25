package main

import (
	"context"
	"english_bot/database"
	"english_bot/handlers"
	handlerUseCase "english_bot/internal/messageHandler/usecase"
	updateUseCase "english_bot/internal/updates/usecase"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURI := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	tgToken := os.Getenv("TG_TOKEN")

	ctx := context.Background()

	//TODO: add AE_KEY and encrypt config
	//var shield = secure.NewShield(os.Getenv("AE_KEY"))

	// viperInstance, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	// }

	// cfg, err := config.ParseConfig(viperInstance)
	// if err != nil {
	// 	log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	// }

	//TODO: add AE_KEY and encrypt config
	//config.DecryptConfig(cfg, shield)

	//logger := logger.NewLogger(cfg).Sugar()
	//defer logger.Sync()

	//TODO: вынести в отдельный файл инициализацию бд
	db, err := database.Connect(ctx, dbName, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	userCollection, err := db.Collection("users")
	//if err != nil {
	//	log.Fatal(err)
	//}
	taskCollection, err := db.Collection("tasks")

	//typeCollection, err := db.Collection("task_types")
	//if err != nil {
	//	log.Fatal(err)
	//}
	progressCollection, err := db.Collection("users_progress")

	userRepo := database.NewUserRepository(userCollection)
	taskRepo := database.NewTaskRepository(taskCollection)
	//typeRepo := database.NewTypeRepository(typeCollection)
	progressRepo := database.NewProgressRepository(progressCollection)

	bot, err := tgbotapi.NewBotAPI(tgToken) // cfg.Telegram.Token
	if err != nil {
		log.Fatal(err)
	}
	updateUC := updateUseCase.Init(bot)

	progressHandler := handlers.NewProgressHandle(taskRepo, userRepo, progressRepo)
	handler := handlerUseCase.InitHandler(bot, userRepo, progressHandler)

	//updates, _ := updateUC.NewUpdates()

	updates, err := bot.GetUpdatesChan(updateUC.UpdateConfig)
	//err = admin.HandleTasks(ctx)
	if err != nil {
		return
	}
	err = handler.HandleMessages(ctx, updates)
	if err != nil {
		return
	}

}
