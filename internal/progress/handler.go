package progress

import (
	"context"
	"english_bot/internal/progress/repository"
	repository2 "english_bot/internal/task/repository"
	progress2 "english_bot/internal/user/repository"
	"english_bot/models"
	"github.com/google/uuid"
	"log"
)

type ProgressHandler struct {
	taskRepo     *repository2.TaskRepository
	userRepo     *progress2.UserRepository
	progressRepo *repository.ProgressRepository
}

func NewProgressHandler(tr *repository2.TaskRepository, ur *progress2.UserRepository, pr *repository.ProgressRepository) *ProgressHandler {
	return &ProgressHandler{
		taskRepo:     tr,
		userRepo:     ur,
		progressRepo: pr,
	}
}

func (p *ProgressHandler) CheckUserAnswer(ctx context.Context, rightAnswer string, userAnswer string, userId int, taskId uuid.UUID) (string, error) {
	var (
		score uint8
		msg   string
	)
	if rightAnswer == userAnswer {
		score = 2
		msg = "верно ☺️"
	} else {
		score = 0
		msg = "неверно 🙁"
	}
	log.Println("inserting user result...", score)

	oldProgress, err := p.progressRepo.GetUserProgress(ctx, userId, taskId)
	if err != nil {
		log.Println("error getting old user progress")
		return "", err
	}
	var (
		oldScore      int64  = 0
		oldTaskLevel  string = task.Level
		receivedTasks int16  = 0
	)
	if oldProgress != nil {
		oldScore = oldProgress.Score
		oldTaskLevel = oldProgress.TaskLevel
		receivedTasks = oldProgress.ReceivedTasks
	}
	res := &models.UserProgress{
		UserID:        userId,
		Score:         oldScore + int64(score),
		TaskLevel:     oldTaskLevel,
		ReceivedTasks: receivedTasks + 1,
	}

	err = p.progressRepo.InsertUserResult(ctx, res)
	if err != nil {
		return "", err
	}
	return msg, nil
}
