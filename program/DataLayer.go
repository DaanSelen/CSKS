package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func initDBConnection() {
	entry, _ := sql.Open("sqlite3", "./csksdatabase.db")
	statement1, _ := entry.Prepare("CREATE TABLE IF NOT EXISTS entry (id INTEGER PRIMARY KEY, hostname varchar(100) NOT NULL, comp varchar(100) NOT NULL, time varchar(200) NOT NULL)")
	defer statement1.Close()
	log.Println("TEST")
	statement1.Exec()
	log.Println("TEST")
	data, err := entry.Query("select count(id) from entry")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	data.Next()
	var currentEntries int
	data.Scan(&currentEntries)
	fmt.Println("Current entries:", currentEntries)
}
