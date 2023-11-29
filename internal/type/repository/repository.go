package repository

import (
	"context"
	"english_bot/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TypeRepository struct {
	collection *mongo.Collection
}

func NewTypeRepository(collection *mongo.Collection) *TypeRepository {
	return &TypeRepository{
		collection: collection,
	}
}

func (tr *TypeRepository) AddType(ctx context.Context, taskType *models.TaskType) error {
	_, err := tr.collection.InsertOne(ctx, taskType)
	if err != nil {
		return fmt.Errorf("error while adding task: %w", err)
	}
	log.Println("task added")
	return nil
}

func (tr *TypeRepository) GetType(typeID int) (*models.TaskType, error) {
	filter := bson.M{"type_id": typeID}
	var taskType models.TaskType
	err := tr.collection.FindOne(context.Background(), filter).Decode(&taskType)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &taskType, err
}
