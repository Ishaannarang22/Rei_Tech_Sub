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

// Create new User
func CreateUser(userIn User) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	insertUserSQL := `INSERT INTO users(username, password, token) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec(userIn.Username, userIn.Password, userIn.Token)
	if err != nil {
		return err
	}

	return nil
}

// Read the user data
func ReadUser(userIn User) (User, error) {
	var err error

	db, err := sql.Open("sqlite3", "./users.db")

	var userOut User

	query := "SELECT password, token FROM users WHERE username=?"
	err = db.QueryRow(query, userIn.Username).Scan(&userOut.Password, &userOut.Token)

	db.Close()

	return userOut, err
}
