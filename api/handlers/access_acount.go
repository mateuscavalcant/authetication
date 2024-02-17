package handlers

import (
	"authentication/api/utils"
	CON "authentication/pkg/database"
	"authentication/pkg/models"
	"authentication/pkg/models/err"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AccessUserAccount handles user authentication.
func AccessUserAccount(c *gin.Context) {
	// Define a struct to hold user information.
	var user models.AccessUser

	// Extract email and password from the POST form data.
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))

	// Create a response object to handle errors.
	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	// Get the database connection.
	db := CON.DB()

	// Query the database to retrieve user information.
	err := db.QueryRow("SELECT id, email, password FROM user WHERE email=?", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		// Log the error and set appropriate response for invalid credentials.
		log.Println("Error executing SQL statement:", err)
		resp.Error["credentials"] = "Invalid credentials"
	}

	// Compare the provided password with the hashed password retrieved from the database.
	encErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if encErr != nil {
		// Set appropriate response for invalid password.
		resp.Error["password"] = "Invalid password"
	}

	// If there are errors, return the response.
	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	// If authentication is successful, store user information in session and return success message.
	session := utils.GetSession(c)
	session.Values["id"] = strconv.Itoa(user.ID)
	session.Values["email"] = user.Email
	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{"message": "User logged in successfully"})
}