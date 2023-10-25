package handlers

import (
	"context"
	"english_bot/database"
	"english_bot/models"
)

type ProgressHandle struct {
	taskRepo     *database.TaskRepository
	userRepo     *database.UserRepository
	progressRepo *database.ProgressRepository
}

func NewProgressHandle(tr *database.TaskRepository, ur *database.UserRepository, pr *database.ProgressRepository) *ProgressHandle {
	return &ProgressHandle{
		taskRepo:     tr,
		userRepo:     ur,
		progressRepo: pr,
	}
}

func (ph *ProgressHandle) GetExercise(ctx context.Context, userID int, typeId int) (*models.Task, error) {
	var (
		taskReceived *models.Task
		err          error
		user         *models.User
	)

	user, err = ph.userRepo.UserByID(userID)
	if err != nil {
		return nil, err
	}
	taskReceived, err = ph.taskRepo.GetTaskByLevelAndType(ctx, user.Level, typeId)
	if err != nil {
		return nil, err
	}
	return taskReceived, nil
}

func (ph *ProgressHandle) GetExerciseFillGaps() error {
	// var (
	// 	typeId int
	// 	level  string
	// )
	// typeId = 2
	// level = "" // get from user information
	// GetTaskByLevelAndType(typeId, level)
	return nil
}

// func ()
