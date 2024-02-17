package models

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name" binding:"required, min=4,max=32"`
	Email           string `json:"email" binding:"required, email"`
	Password        string `json:"password" binding:"required, min=8, max=16"`
	ConfirmPassword string `json:"cpassword" binding:"required"`
}