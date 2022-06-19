package main

import (
	"fmt"
	"log"
	"strconv"
)

type Klant struct {
	ID         int    `json:"id"`
	Voornaam   string `json:"voornaam"`
	Achternaam string `json:"achternaam"`
	Aankomst   string `json:"aankomst"`
	Vertrek    string `json:"vertrek"`
	Staplaats  string `json:"staplaats"`
	Opmerking  string `json:"opmerking"`
}

func main() {
	fmt.Println("Welcome to the CSKS System.")
	go initDBConnection()
	initHTTP()
}

func handleError(err error, location ...string) {
	if err != nil {
		log.Println("Error encountered at", location, "error:", err)
	}
}

func postKlantAdd(voor, achter, aan, ver, sta, opmerk string) {
	klantAdd(voor, achter, aan, ver, sta, opmerk)
}

func deleteRemoveKlant(idQueryStr string) bool {
	idQueryInt, err := strconv.Atoi(idQueryStr)
	if err != nil {
		return false
	} else {
		removeKlant(idQueryInt)
	}
	return true
}

func getSearchData(query string) []Klant {
	return searchData(query)
}

func getKlantAll() []Klant {
	return KlantAll()
}
