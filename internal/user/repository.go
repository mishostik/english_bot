package user

import (
	"context"
)

type Repository interface {
	UserByID(ctx context.Context, userID int) (*User, error)
	RegisterUser(ctx context.Context, user *User) error
}
