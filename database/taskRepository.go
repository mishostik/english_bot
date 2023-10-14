package database

import (
	"context"
	"english_bot/models"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
