package task

import (
	"context"
)

type Repository interface {
	GetTaskByLevelAndType(ctx context.Context, level string, taskType int) (*Task, error)
	GetRandomTaskByLevel(ctx context.Context, level string) (*Task, error)
}
