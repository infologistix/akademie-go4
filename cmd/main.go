package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// erstelle neuen HTTP-Server
	router := http.NewServeMux()

	// mit folgendem Befehl wird ein Handler für eine HTTP-Route hinzugefügt
	router.HandleFunc("/all", getAll)
	router.HandleFunc("/cities", cities)
	router.HandleFunc("/projects", projects)

	// Ab hier beginnt die initialisierung des HTTP-Servers
	// Yohanes war hier !!
	fmt.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
