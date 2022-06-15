package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	customer *sql.DB
)

func initDBConnection() {
	log.Println("DATABASE SERVER INITIALISING")
	customer, _ = sql.Open("sqlite3", "./SQLite3Klant.db")
	stmnt, err := customer.Prepare("CREATE TABLE IF NOT EXISTS customer (id INTEGER PRIMARY KEY, voornaam TEXT NOT NULL, achternaam TEXT NOT NULL, aankomst TEXT NOT NULL, vertrek TEXT NOT NULL, staplaats TEXT NOT NULL, opmerking TEXT)")
	handleError(err, "Preparing the statement CREATE TABLE IF NOT EXISTS...")
	_, err = stmnt.Exec()
	handleError(err, "Executing the statement CREATE TABLE IF NOT EXISTS...")
	log.Println("Current entries:", getCount())
}

func getCount() int {
	data, err := customer.Query("SELECT count(id) FROM customer")
	handleError(err, "Querying SELECT count(id) FROM customer")
	defer data.Close()
	data.Next()
	var currentEntries int
	data.Scan(&currentEntries)
	return currentEntries
}

func klantAdd(voor, achter, aan, ver, sta, opmerk string) {
	statement, _ := customer.Prepare("INSERT INTO customer (id, voornaam, achternaam, aankomst, vertrek, staplaats, opmerking) VALUES (null, ?, ?, ?, ?, ?, ?)")
	defer statement.Close()
	statement.Exec(voor, achter, aan, ver, sta, opmerk)
	log.Println("[KLANTADD] DONE")
}

func removeKlant(idInt int) {
	statement, _ := customer.Prepare("DELETE FROM customer WHERE id = ?")
	defer statement.Close()
	_, err := statement.Exec(idInt)
	handleError(err, "Executing Query DELTE * FROM customer")
	log.Println("[KLANTREMOVE] DONE")
}

func KlantAll() []Klant {
	var klanten []Klant
	rows, err := customer.Query("SELECT * FROM customer")
	handleError(err, "Querying SELECT * FROM customer")
	defer rows.Close()
	for rows.Next() {
		var klant Klant
		rows.Scan(&klant.ID, &klant.Voornaam, &klant.Achternaam, &klant.Aankomst, &klant.Vertrek, &klant.Staplaats, &klant.Opmerking)
		klanten = append(klanten, klant)
	}
	return klanten
}
