package main

import (
	"context"
	"english_bot/internal/core"
	"english_bot/internal/dictionary"
	incorrectRepository "english_bot/internal/incorrect/repository"
	"english_bot/internal/message"
	messageUseCase "english_bot/internal/message/usecase"
	"english_bot/internal/progress/repository"
	stateUseCase "english_bot/internal/state/usecase"
	taskRepository "english_bot/internal/task/repository"
	updateUseCase "english_bot/internal/updates/usecase"
	userRepository "english_bot/internal/user/repository"
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
	//config.DecryptConfig(cfg, shield)

	//logger := logger.NewLogger(cfg).Sugar()
	//defer logger.Sync()

	//TODO: вынести в отдельный файл инициализацию бд
	db, err := core.Connect(ctx, "english_bot", "")
	if err != nil {
		log.Fatal(err)
	}

	userCollection, err := db.Collection("users")
	taskCollection, err := db.Collection("tasks")

	progressCollection, err := db.Collection("users_progress")
	incAnswersCollection, err := db.Collection("incorrect_answers")

	userRepo := userRepository.NewUserRepository(userCollection)
	taskRepo := taskRepository.NewTaskRepository(taskCollection)
	progressRepo := repository.NewProgressRepository(progressCollection)

	incorrectAnswersRepo := incorrectRepository.NewIncorrectRepository(incAnswersCollection)

	bot, err := tgbotapi.NewBotAPI("") // cfg.Telegram.Token
	if err != nil {
		log.Fatal(err)
	}
	updateUC := updateUseCase.Init(bot)
	stateUC := stateUseCase.NewStateUseCase()
	messageUC := messageUseCase.NewMessageHandlerUsecase(userRepo, taskRepo, incorrectAnswersRepo, progressRepo, &stateUC)

	dictionaryHandler := dictionary.NewDictionaryHandler()

	handler := message.InitHandler(bot, &stateUC, dictionaryHandler, messageUC)

	updates, err := bot.GetUpdatesChan(updateUC.UpdateConfig)
	if err != nil {
		return
	}
	err = handler.HandleMessages(ctx, updates)
	if err != nil {
		return
	}

}
