package utilsdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string
	Password string
	Token    string
}

// Read the user data
func ReadUser(userIn User) (User, error) {
	var err error

	db, err := sql.Open("sqlite3", "./userData.db")

	var userOut User

	query := "SELECT password, token FROM users WHERE username=?"
	err = db.QueryRow(query, userIn.Username).Scan(&userOut.Password, &userOut.Token)

	db.Close()

	return userOut, err
}
