package _type

import (
	"context"
)

type Repository interface {
	AddType(ctx context.Context, taskType *TaskType) error
	GetType(typeID int) (*TaskType, error)
}
