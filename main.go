package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "myorm.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()
	_, _ = db.Exec("DROP TABLE IF EXISTS Users;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "toms", "sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err != nil {
		log.Println(name)
	}
}
