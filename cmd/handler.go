package main

// immer nötig!
import (
	"encoding/json"
	"log"
	"net/http"
)

// initiieren einer Liste mit Customer Inhalt
var l_customers = []Customer{
	{Name: "AWS", Address: "München", Tel: "+49 000 12345678", Email: "info@aws.de"}, // definite Typangabe
	{"infologistix GmbH", "Garching", "+49 111 87654321", "info@infologistix.de"},    // Reihenfolge basierende Typanagabe
	{"NordLB", "Hannover", "123", "@nordlb.de"},
}

var l_cities = []City{
	{Name: "Hamburg", Einwohner: 2e6, Bundesland: "HH"},
	{Name: "Aachen", Einwohner: 200000, Bundesland: "NRW"},
	{Name: "Frankfurt", Einwohner: 700000, Bundesland: "H"},
	{Name: "Kaiserslautern", Einwohner: 99845, Bundesland: "RP"},
	{"Hannover", 200000, "NI"},
}

var l_projects = []Project{
	{Name: "BA", Customer: l_customers[1], City: l_cities[0], Budget: "300.000"},
	{"SPOT", l_customers[3], 4, "123.456"},
}

var l_members = []InfologistixMembers{
	{Name: "Paul", City: l_cities[1], Email: "paul.schmidt@infologistix.de", Project: l_projects[0]},
	{Name: "Yohanes", City: l_cities[3], Email: "yohanes.korbafo@infologistix.de", Project: l_projects[1]},
}

// function getAll
// wird mit einem Handler aus ResponseWriter und dazugehörigem Pointer Request aufgebaut
// die Parameter werden der Funktion als w & r übergeben
func getAll(w http.ResponseWriter, r *http.Request) {
	// hier könnte noch eine Abfrage gegen den Request laufen (z.B URL, URL Parameter)
	// diese "if" Abfrage testet, ob l_customers erfolgreich und ohne Fehler durchläuft
	// ...Encode(data) liefert einen Fehler 'error' wenn etwas schief läuft
	if json.NewEncoder(w).Encode(l_customers) != nil {
		log.Fatal(json.NewEncoder(w).Encode(l_customers)) // logge den Fehler direkt
		log.Fatalf("There was some issue")                // Logge etwas als String
	}
}

// Hier ist Platz für weitere Handler
func cities(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.RequestURI)
	if json.NewEncoder(w).Encode(l_cities) != nil {
		log.Fatal(json.NewEncoder(w).Encode(l_cities)) // logge den Fehler direkt
		log.Fatalf("There was some issue")             // Logge etwas als String
	}
}

// noch eine Funktion
func projects(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.RequestURI)
	if json.NewEncoder(w).Encode(l_projects) != nil {
		log.Fatal(json.NewEncoder(w).Encode(l_projects)) // logge den Fehler direkt
		log.Fatalf("There was some issue")               // Logge etwas als String
	}
}
