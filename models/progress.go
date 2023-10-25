package models

type Progress struct {
	UserID   int    `bson:"task_id"`
	TypeID   int    `bson:"type_id"`
	Level    string `bson:"level"`
	Question string `bson:"question"`
	Answer   string `bson:"answer"`
}

//
//type TaskType struct {
//	TypeID int    `bson:"type_id"`
//	Type   string `bson:"type"`
//}
