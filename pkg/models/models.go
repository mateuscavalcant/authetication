package models

type CreateUser struct {
	ID              int    `json:"id"`
	Name            string `json:"name" binding:"required, min=4,max=32"`
	Email           string `json:"email" binding:"required, email"`
	Password        string `json:"password" binding:"required, min=8, max=16"`
	ConfirmPassword string `json:"cpassword" binding:"required"`
}

type AccessUser struct {
	ID              int    `json:"id"`
	Email           string `json:"email" binding:"required, email"`
	Password        string `json:"password" binding:"required, min=8, max=16"`
	
}

type UpdateUserAccount struct {
	Icon []byte
	Name            string `json:"name" binding:"required, min=4,max=32"`
	Bio            string `json:"bio" binding:"required, max=150"`
	Email           string `json:"email" binding:"required, email"`
	Password        string `json:"password" binding:"required, min=8, max=16"`
}