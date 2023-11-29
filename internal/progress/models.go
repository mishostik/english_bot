package progress

type UserProgress struct {
	UserID int `bson:"user_id"`
	//UserScore     uint8 `bson:"user_score"`    // 0, 1 или 2 балла за каждый ответ
	TaskLevel     string `bson:"task_level"`     // A0, A1, B1, B2, C1
	ReceivedTasks int16  `bson:"received_tasks"` // received tasks for certain level
	Score         int64  `bson:"score"`
}
