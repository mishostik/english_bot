package database

import (
	"context"
	"english_bot/models"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		collection: collection,
	}
}

func (tr *TaskRepository) AddTask(ctx context.Context, task *models.Task) error {
	_, err := tr.collection.InsertOne(ctx, task)
	if err != nil {
		return fmt.Errorf("error while adding task: %w", err)
	}
	log.Println("task added")
	return nil
}

func (tr *TaskRepository) RandomTask(level string) {

}

func (tr *TaskRepository) GetTaskByLevelAndType(ctx context.Context, level string, taskType int) (*models.Task, error) {

	filter := bson.M{
		"type_id": taskType,
		"level":   level,
	}

	var result models.Task
	err := tr.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("task not found")
		} else {
			log.Println("some error")
		}
	} else {
		return &result, nil
	}
	return nil, err
}
