package usecase

import (
	"context"
	"english_bot/constants"
	"english_bot/database"
	"english_bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TaskHandler struct {
	bot      *tgbotapi.BotAPI
	taskRepo *database.TaskRepository
	typeRepo *database.TypeRepository
}

func InitAdmin(bot *tgbotapi.BotAPI, taskRepo *database.TaskRepository, typeRepo *database.TypeRepository) *TaskHandler {
	return &TaskHandler{
		bot:      bot,
		taskRepo: taskRepo,
		typeRepo: typeRepo,
	}
}

func (h *TaskHandler) HandleTasks(ctx context.Context) error {
	taskType, err := h.typeRepo.GetType(1)
	if err != nil {
		log.Println(err)
	}
	task := models.Task{
		TaskID:   1,
		TypeID:   taskType.TypeID,
		Level:    constants.LevelA1,
		Question: "Вопрос",
		Answer:   "Ответ",
	}
	err = h.taskRepo.AddTask(ctx, task)
	if err != nil {
		log.Println(err)
	}
	return nil
}
