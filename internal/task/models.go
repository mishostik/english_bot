package task

import "github.com/google/uuid"

type Task struct {
	ID       uuid.UUID `bson:"task_id"`
	TypeID   int       `bson:"type_id"`
	Level    string    `bson:"level"`
	Question string    `bson:"question"`
	Answer   string    `bson:"answer"`
}
