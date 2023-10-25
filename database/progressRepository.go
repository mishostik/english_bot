package database

import "go.mongodb.org/mongo-driver/mongo"

type ProgressRepository struct {
	collection *mongo.Collection
}

func NewProgressRepository(collection *mongo.Collection) *ProgressRepository {
	return &ProgressRepository{
		collection: collection,
	}
}

func (pr *ProgressRepository) CalculateUserAnswerScore() {
	// received answer equal to answer from collection? => 2, 1 or 0
}

func (pr *ProgressRepository) CheckAnswer() (int, error) {
	// get info from user
	return 0, nil
}
