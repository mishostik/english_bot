package user

import "time"

type User struct {
	UserID       int       `bson:"user_id"`
	Username     string    `bson:"username"`
	RegisteredAt time.Time `bson:"registered_at"`
	LastActiveAt time.Time `bson:"last_active_at"`
	Level        string    `bson:"level"`
}
