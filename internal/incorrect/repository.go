package incorrect

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	GetAnswers(ctx context.Context, taskId uuid.UUID) (*Answers, error)
}
