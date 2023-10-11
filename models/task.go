package models

type Task struct {
	TaskID   int    `bson:"task_id"`
	Type     string `bson:"type"`
	Level    string `bson:"level"`
	Question string `bson:"question"`
	Answer   string `bson:"answer"`
}
