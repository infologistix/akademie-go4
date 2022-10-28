package main

// <- Dies sind Kommentar Angaben
// Dies ist ein Sample
// Die einzelnen Structs sind selber hier anzulegen. Der Merge der Daten sollte ohne Probleme vollzogen werden können
// Grundsätzlich sollte der Code auch dokumentiert werden. Wie auch hier gilt, je mehr Inhalt
// in weniger Zeilen ist hilfreicher, als ein ganzer Dokublock.
// Hinweis: Die Einrückungen dienen lediglich der Übersicht!
type Customer struct {
	Name    string `json:"name"`
	Address string `json:"adress"`
	Tel     string `json:"tel"`
	Email   string `json:"mail"`
}

// Ab Hier können einzelne Structs hinzugefpügt werden
type City struct {
	Name       string `json:"name"`
	Einwohner  int    `json:"einwohner"`
	Bundesland string `json:"bundesland"`
}

type Project struct {
	Name     string   `json:"name"`
	Customer Customer `json:"customer"`
	City     City     `json:"city"`
	Budget   string   `json:"budget"`
}

type InfologistixMembers struct {
	Name    string  `json:"name"`
	Email   string  `json:"mail"`
	City    City    `json:"city"`
	Project Project `json:"project"`
}

//Neue Zeile
