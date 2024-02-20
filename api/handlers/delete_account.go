package handlers

import (
	"authentication/api/utils"
	"authentication/pkg/database"
	"authentication/pkg/models"
	"authentication/pkg/models/err"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUserAccount(c *gin.Context) {
	idInterface, _ := utils.AllSessions(c)
	id, _ := strconv.Atoi(idInterface.(string))

	// Define a struct to hold user information.
	var user models.User

	// Extract email and password from the POST form data.
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	confirmPassword := strings.TrimSpace(c.PostForm("confirm_password"))

	// Create a response object to handle errors.
	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	// Get the database connection from the pool.
	db := database.GetDB()

	// Query the database to retrieve user information.
	row := db.QueryRow("SELECT id, email, password FROM user WHERE email=?", email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		// Log the error and set appropriate response for invalid credentials.
		log.Println("Error executing SQL statement:", err)
		resp.Error["email"] = "Invalid credentials"
	}

	// Compare the provided password with the hashed password retrieved from the database.
	encErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if encErr != nil || password != confirmPassword {
		// Set appropriate response for invalid password.
		resp.Error["password"] = "Invalid password"
	}

	// If there are errors, return the response.
	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	// If authentication is successful, delete user.
	stmt, errDB := db.Prepare("DELETE FROM user WHERE id=?")
	if errDB != nil {
		log.Println("Error preparing SQL statement:", errDB)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	defer stmt.Close()

	_, errDB = stmt.Exec(id)

	c.JSON(200, gin.H{"message": "User deleted in successfully"})

	// Close the prepared statement to release resources.
	defer stmt.Close()

}