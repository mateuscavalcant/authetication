package validators

import (
	CON "authetication/pkg/database"

	"github.com/badoux/checkmail"
)


func ValidateFormatEmail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}
	return nil
}

func ExistEmail(email string) (bool, error) {
	db := CON.DB()

	var emailCount int

	err := db.QueryRow("SELECT COUNT(id) AS emailCount FROM user1 WHERE email=?", email).Scan(&emailCount)
	if err != nil {
		return false, err
	}

	return emailCount > 0, nil
}