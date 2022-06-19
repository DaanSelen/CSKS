package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initHTTP() {
	log.Println("API SERVER INITIALISING")
	csks := mux.NewRouter().StrictSlash(true)

	csks.HandleFunc("/", handleRootEndpoint)
	csks.HandleFunc("/klant", handleKlantEndpoint).Methods("GET")
	csks.HandleFunc("/klant/add", handleKlantAddEndpoint).Methods("POST")
	csks.HandleFunc("/klant/remove", handleKlantRemoveEndpoint).Methods("DELETE")

	http.ListenAndServe((":1269"), csks)
}

func handleRootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	json.NewEncoder(w).Encode("CeldrithService KlantSysteem v0.1, Root Endpoint.")
}

func handleKlantAddEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var klant Klant
	_ = json.NewDecoder(r.Body).Decode(&klant)
	postKlantAdd(klant.Voornaam, klant.Achternaam, klant.Aankomst, klant.Vertrek, klant.Staplaats, klant.Opmerking)
	json.NewEncoder(w).Encode("Adding customer to database.")
}

func handleKlantRemoveEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idQuery, ok := r.URL.Query()["id"]
	if ok && len(idQuery) > 0 && deleteRemoveKlant(idQuery[0]) {
		json.NewEncoder(w).Encode("Removing customer from database")
	} else {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Missing or incorrect URL query 'id'.")
	}
}

func handleKlantEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	searchQuery, ok := r.URL.Query()["search"]
	if ok && len(searchQuery) > 0 {
		result := getSearchData(searchQuery[0])
		if result == nil {
			json.NewEncoder(w).Encode("No results found")
		} else {
			json.NewEncoder(w).Encode(result)
		}
	} else {
		klanten := getKlantAll()
		if klanten == nil {
			json.NewEncoder(w).Encode("Empty Database detected. No information present.")
		} else {
			json.NewEncoder(w).Encode(klanten)
		}
	}
}
