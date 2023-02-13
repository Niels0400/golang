package main

import (
	"fmt"
	"time"
)

func main() {
	// Hardcoded reeks van kentekens
	allowedPlates := [5]string{"Kenteken1", "Kenteken2", "Kenteken3", "Kenteken4", "Kenteken5"}

	// Huidige tijd
	currentTime := time.Now()

	// Kenteken opvragen
	var kenteken string
	fmt.Print("Voer jou kenteken in: ")
	fmt.Scanln(&kenteken)

	// Controleren of het kenteken in de reeks voorkomt
	plateAllowed := false
	for _, p := range allowedPlates {
		if kenteken == p {
			plateAllowed = true
			break
		}
	}

	// Bericht afhankelijk van kenteken en tijd
	if !plateAllowed {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	} else if currentTime.Hour() >= 7 && currentTime.Hour() < 12 {
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken")
	} else if currentTime.Hour() >= 12 && currentTime.Hour() < 18 {
		fmt.Println("Goedemiddag! Welkom bij Fonteyn Vakantieparken")
	} else if currentTime.Hour() >= 18 && currentTime.Hour() < 23 {
		fmt.Println("Goedenavond! Welkom bij Fonteyn Vakantieparken")
	} else {
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten")
	}
}
