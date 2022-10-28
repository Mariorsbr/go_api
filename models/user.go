package models


type User struct{
	ID     string  `json:"id" gorm:"primary_key"`
    Username  string  `json:"username"`
    Password string  `json:"password"`
}

type CreateUserInput struct {
	Username  string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Username  string  `json:"username"`
	Password string  `json:"password"`
}


