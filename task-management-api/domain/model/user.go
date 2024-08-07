package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	ID              primitive.ObjectID `bson:"_id,omitempty"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
