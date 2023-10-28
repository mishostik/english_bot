package main

import (
	"context"
	"english_bot/database"
	"english_bot/handlers"
	handlerUseCase "english_bot/internal/messageHandler/usecase"
	updateUseCase "english_bot/internal/updates/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	//dbURI := os.Getenv("DB_URI")
	//dbName := os.Getenv("DB_NAME")
	//tgToken := os.Getenv("TG_TOKEN")

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
	//db, err := database.Connect(ctx, dbName, dbURI)
	db, err := database.Connect(ctx, "", "")
	if err != nil {
		log.Fatal(err)
	}

	userCollection, err := db.Collection("users")
	taskCollection, err := db.Collection("tasks")

	progressCollection, err := db.Collection("users_progress")

	userRepo := database.NewUserRepository(userCollection)
	taskRepo := database.NewTaskRepository(taskCollection)
	//typeRepo := database.NewTypeRepository(typeCollection)
	progressRepo := database.NewProgressRepository(progressCollection)

	bot, err := tgbotapi.NewBotAPI("") // cfg.Telegram.Token
	if err != nil {
		log.Fatal(err)
	}
	updateUC := updateUseCase.Init(bot)

	messageUsecase := handlerUseCase.NewMessageHandlerUsecase(userRepo, taskRepo)

	dictionaryHandler := handlers.NewDictionaryHandler()
	exerciseHandler := handlers.NewExerciseHandler(messageUsecase)
	progressHandler := handlers.NewProgressHandler(taskRepo, userRepo, progressRepo)

	generalHandler := handlers.NewGeneralHandler(dictionaryHandler, exerciseHandler, progressHandler, messageUsecase, userRepo)

	handler := handlerUseCase.InitHandler(bot, generalHandler)

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
