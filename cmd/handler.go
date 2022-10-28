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
