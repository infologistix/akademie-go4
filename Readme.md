# Team Hackathon infologistix

Dieses Repository dient als Team Hackathon Startpunkt einer Freitagsakademie.

## Benutzung
Git Bash:

```bash
$ go run cmd/*
```
Powershell:
```bash
$ cd cmd
$ go run .
```
Mit dem Befehl lässt sich der Webserver starten und man kann unter der im Fenster angezeigten URL auf den eigenen geschriebenen Dienst zugreifen.

## GO-Strukturen

In Go werden die Pakete und Funktionen hierarisch und in logischen Gruppen in einzelnen Dateien zusammengefasst. Dabei gilt, dass die entsprechenden Dateien Hinweis darauf geben, welchen Sinn der Inhalt besitzt.

In diesem Fall sieht die Struktur wie folgt aus:
```
|
| - go.mod # beschreibt die Go-Module, die benötigt werden und das Package
| - cmd
|   | - main.go
|   | - data.go
|   | - handler.go
```
Diese Strukturen und Benamungen sind falls möglich beizubehalten, wenngleich für jeden Handler (HTTP-Endpunkt) eine eigene Datei erstellt werden kann. Die Verlinkung innerhalb der `main.go` erfolgt automatisch.

### main.go

Die Hauptdatei, in welcher der Rest-Dienst gestartet wird. Sämtliche Handler werden hier registriert und werden verfügbar gemacht.

### data.go

Auslagerung der Datentypen. Diese können hier im `Struct` Format abgelegt werden. Es erfolgt neben der Typisierung (zb: `string`) auch die JSON-Notation und deren Beschreibung. In der Datei ist auch ein Beispiel für einen Datensatz zu finden. 

### handler.go

Die eigentliche Magie entsteht hier an dieser Stelle durch hinzufügen von spezifischen Abläufen für bestimmte HTTP-Routen. Diese können weiter spezifiziert werden, als das Beispiel zeigt.

## Ziel
Ziel ist das Erproben von Git in der Praxis. Dafür sollen alle aus dem Team folgende Operationen mindestens einmal durchgeführt haben:
<ul>
    <li>Eins der zur Verfügung stehenden Repositories klonen</li>
    <li>Eine Änderung im Code pushen, die zu einer funktionalen Veränderung des Codes beiträgt (also nicht nur einen einzeiligen Kommentar)</li>
    <ul>
        <li>Was sind die einzelnen Schritte?</li>
    </ul>
    <li>Einen Pull durchführen, nachdem jemand anderes etwas ins Repo gepusht hat</li>
    <ul>
        <li>Wo stehen die letztes Commits?</li>
        <li>Wer hat die letzte Änderung vorgenommen und was wurde geändert?</li>
    </ul>
    <li>Einen Mergekonflikt lösen</li>
    <ul>
        <li>Was ruft einen Mergekonflikt hervor?</li>
        <li>Welche Möglichkeiten gibt es einen Mergekonflikt zu lösen?</li>
    </ul>
    <li>Welche Befehle gibt es noch (https://git-scm.com/docs)? Probiert weitere aus.</li>
    <ul>
        <li>Sucht euch einen heraus und versucht zu beschreiben, was man damit machen kann.</li>
    </ul>
</ul>

Die Antworten auf die Fragen können gerne direkt hinter die Fragen geschrieben und anschließend gepusht werden.