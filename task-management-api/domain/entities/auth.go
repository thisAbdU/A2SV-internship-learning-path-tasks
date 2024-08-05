package entities

import "example/GO-PRACTICE-EXERCISE/GO-API-exercise/domain/model"

type AuthenticatedUser struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Token struct {
	Token string `json:"token"`
}


type AuthUseCase interface {
	Register(currUser AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error)
	Login(currUser AuthenticatedUser, userLogin *model.UserLogin, param any) (Token, string, error)
	AdminRegister(currUser AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error)
}
