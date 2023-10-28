package handlers

import (
	"english_bot/database"
)

type ProgressHandler struct {
	taskRepo     *database.TaskRepository
	userRepo     *database.UserRepository
	progressRepo *database.ProgressRepository
}

func NewProgressHandler(tr *database.TaskRepository, ur *database.UserRepository, pr *database.ProgressRepository) *ProgressHandler {
	return &ProgressHandler{
		taskRepo:     tr,
		userRepo:     ur,
		progressRepo: pr,
	}
}
