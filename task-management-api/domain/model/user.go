package model

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Id string `json:"id"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id 	 string `json:"id"`
}

type UserInfo struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
