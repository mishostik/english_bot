package database

import (
	"context"
	"english_bot/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// fucking fuck
//type UserRepository interface {
//	UserByID(userID int64) (*models.User, error)
//	RegisterUser(user *models.User) (bool, error)
//}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) UserByID(userID int64) (*models.User, error) {
	//userCollection := ur.collection.Collection("users") // double shit
	filter := bson.M{"user_id": userID}
	var user models.User
	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
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
