package models

type Task struct {
	TypeID   int    `bson:"type_id"`
	Level    string `bson:"level"`
	Question string `bson:"question"`
	Answer   string `bson:"answer"`
}

type TaskType struct {
	TypeID int    `bson:"type_id"`
	Type   string `bson:"type"`
}

type CurrentTask struct {
	UserID        int    // ID пользователя
	CurrentTaskID int    // ID текущего задания
	CorrectAnswer string // Правильный ответ
}
