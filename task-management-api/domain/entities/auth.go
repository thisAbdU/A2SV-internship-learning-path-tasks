package entities

import "task-management-api/domain/model"

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
	Register(userCreate *model.UserCreate) (*model.UserInfo,error)
	Login(userLogin *model.UserLogin) (string,error)
	AdminRegister(currUser AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo,error)
}
