package handlers

import (
	CON "authetication/pkg/database"
	"authetication/pkg/models"
	"authetication/pkg/models/err"
	"authetication/pkg/validators"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUserAccount(c *gin.Context) {
	var user models.User

	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	name := strings.TrimSpace(c.PostForm("name"))
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	confirmPassword := strings.TrimSpace(c.PostForm("confirm_password"))

	existEmail, err := validators.ExistEmail(email)
	if err != nil {
		log.Println("Error checking email existence:", err)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}
	

	if name == "" || email == "" || password == "" || confirmPassword == "" {
		resp.Error["missing"] = "Some values are missing!"
	}

	if len(name) < 4 || len(name) > 32 {
		resp.Error["name"] = "name should be between 4 and 32"
	}

	if validators.ValidateFormatEmail(email) != nil {
		resp.Error["email"] = "Invalid email format!"
	}

	if existEmail {
		resp.Error["email"] = "Email already exists!"
	}
	

	if len(password) < 8 || len(password) > 16 {
		resp.Error["password"] = "Passwords should be between 8 and 16"
	}

	if password != confirmPassword {
		resp.Error["confirm_password"] = "Passwords don't match"
	}

	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	user.Name = name
	user.Email = email
	user.Password = password

	db := CON.DB()

	query := "INSERT INTO user1 (name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Name, user.Email, validators.Hash(user.Password))
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}