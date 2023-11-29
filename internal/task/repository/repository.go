package repository

import (
	"context"
	"english_bot/models"
	"log"
	"math/rand"
	"time"

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

func (r *TaskRepository) GetTaskByLevelAndType(ctx context.Context, level string, taskType int) (*models.Task, error) {
	filter := bson.M{
		"type_id": taskType,
		"level":   level,
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
	}(cursor, ctx)

	var tasks []*models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if len(tasks) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)

	selectedTask := tasks[randomGenerator.Intn(len(tasks))]

	return selectedTask, nil
}

func (r *TaskRepository) GetRandomTaskByLevel(ctx context.Context, level string) (*models.Task, error) {
	log.Println("...getting random task")
	filter := bson.M{
		"level": level,
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal("cursor -", err.Error())
		}
	}(cursor, ctx)

	var tasks []*models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("decode error")
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if len(tasks) == 0 {
		log.Println("error no documents")
		return nil, mongo.ErrNoDocuments
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	log.Println("amount of tasks -", len(tasks))
	selectedTask := tasks[randomGenerator.Intn(len(tasks))]

	return selectedTask, nil
}
