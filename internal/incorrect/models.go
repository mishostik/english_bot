package incorrect

import "github.com/google/uuid"

type Answers struct {
	TaskId uuid.UUID `bson:"task_id"`
	A      string    `bson:"a"`
	B      string    `bson:"b"`
	C      string    `bson:"c"`
}
