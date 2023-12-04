package repository

import (
	"context"
	"english_bot/internal/incorrect"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IncorrectRepository struct {
	collection *mongo.Collection
}

func NewIncorrectRepository(collection *mongo.Collection) *IncorrectRepository {
	return &IncorrectRepository{
		collection: collection,
	}
}

func (r *IncorrectRepository) GetAnswers(ctx context.Context, taskId uuid.UUID) (*incorrect.Answers, error) {
	filter := bson.M{"task_id": taskId}
	var answers incorrect.Answers
	err := r.collection.FindOne(ctx, filter).Decode(&answers)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &incorrect.Answers{
				TaskId: taskId,
				A:      "empty",
				B:      "empty",
				C:      "empty",
			}, nil
		}
		return nil, err
	}
	return &answers, nil
}
