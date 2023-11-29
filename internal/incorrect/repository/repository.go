package repository

import (
	"context"
	"english_bot/models"
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

func (r *IncorrectRepository) GetAnswers(ctx context.Context, taskId uuid.UUID) (*models.IncorrectAnswers, error) {
	filter := bson.M{"task_id": taskId}
	var answers models.IncorrectAnswers
	err := r.collection.FindOne(ctx, filter).Decode(&answers)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &models.IncorrectAnswers{
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
