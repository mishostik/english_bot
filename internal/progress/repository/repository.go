package repository

import (
	"context"
	"english_bot/internal/progress"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ProgressRepository struct {
	collection *mongo.Collection
}

func NewProgressRepository(collection *mongo.Collection) *ProgressRepository {
	return &ProgressRepository{
		collection: collection,
	}
}

func (r *ProgressRepository) InsertUserResult(ctx context.Context, progress *progress.UserProgress) error {
	//_, err := r.collection.InsertOne(ctx, progress)
	//if err != nil {
	//	return fmt.Errorf("error while registering user: %w", err)
	//}
	log.Println("progress updated")
	return nil
}

func (r *ProgressRepository) GetUserProgress(ctx context.Context, userId int, taskId uuid.UUID) (*progress.UserProgress, error) {
	filter := bson.M{
		"task_id": taskId,
		"user_id": userId,
	}
	var response progress.UserProgress
	err := r.collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("error no documents")
			return nil, nil
		}
		return nil, err
	}
	return &response, err
}
