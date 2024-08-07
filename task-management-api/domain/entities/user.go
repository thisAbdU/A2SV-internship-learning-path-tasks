package entities

import (
	"context"
	"task-management-api/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetUser(ctx context.Context, param string) ([]*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, updatedUser User) error
	DeleteUser(ctx context.Context, id string) error
	CreateUser(ctx context.Context, newUser model.UserCreate) (*model.UserInfo, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, updatedUser User) error
	DeleteUser(ctx context.Context, id string) error
	CreateUser(ctx context.Context, newUser User) error
}
