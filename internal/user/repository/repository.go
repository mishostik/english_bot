package repository

import (
	"context"
	"english_bot/models"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) UserByID(ctx context.Context, userID int) (*models.User, error) {
	log.Println("getting user by id", userID)
	filter := bson.M{"user_id": userID}
	var user models.User
	err := ur.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("error no documents")
			return nil, nil
		}
		return nil, err
	}
	log.Println("user exist")
	return &user, err
}

func (ur *UserRepository) RegisterUser(ctx context.Context, user *models.User) error {
	_, err := ur.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("error while registering user: %w", err)
	}
	log.Println("register user")
	return nil
}
