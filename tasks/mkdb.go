package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("../users.db")

	log.Println("Creating DB...")
	file, err := os.Create("../users.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("DB created")

	articlesDB, _ := sql.Open("sqlite3", "../users.db")
	defer articlesDB.Close()
	createTable(articlesDB)

	testValue(articlesDB, "test", "$2a$10$5LxAkDSCFubhQeh4cROSWe.eqeldcpARnHGpIJiM9kAYf8XFaee4a", "455645645641")
}

func createTable(db *sql.DB) {
	articlesTableSQL := `CREATE TABLE users (
		"username" TEXT,
		"password" TEXT,
		"token" TEXT
	);`

	log.Println("Creating Table for Users")
	stmnt, err := db.Prepare(articlesTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmnt.Exec()
	log.Println("Users Table Created")
}

func testValue(db *sql.DB, username, password, token string) {
	log.Println("Adding test value...")
	insertTestValueSQL := `INSERT INTO users (username, password, token) VALUES (?, ?, ?)`
	stmt, err := db.Prepare(insertTestValueSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(username, password, token)
	if err != nil {
		log.Fatal(err.Error())
	}
}
