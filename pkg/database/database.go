package databse

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DB returns a connection to the MySQL database.
func DB() *sql.DB {
	// Retrieve database credentials from environment variables.
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	_db := os.Getenv("DB")

	// Open a connection to the MySQL database.
	db, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:3306)/"+_db)
	if err != nil {
		// Panic if there's an error opening the connection.
		panic(err)
	}

	// Ping the database to ensure the connection is successful.
	err = db.Ping()
	if err != nil {
		// Panic if there's an error pinging the database.
		panic(err)
	}

	// Return the database connection.
	return db
}
