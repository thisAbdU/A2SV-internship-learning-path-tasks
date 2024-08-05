package entities

import (
	"context"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/domain/model"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetUser(ctx context.Context, param string) ([]*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, updatedUser User) error
	DeleteUser(ctx context.Context, id string) error
	CreateUser(ctx context.Context, newUser User) (*model.UserInfo, error)
}

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, updatedUser User) error
	DeleteUser(ctx context.Context, id string) error
	CreateUser(ctx context.Context, newUser User) error
}
