package models

type User struct {
	Id       int64  `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}