package databse

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {

	user := "user"
	password := "passowrd"
	_db := "db"

	db, _ := sql.Open("mysql", user+":"+password+"@tcp(localhost:3306)/"+_db)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}